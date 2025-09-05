package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lutefd/md-publisher/api/internal/api"
	"github.com/lutefd/md-publisher/api/internal/storage"
)

func TestServerSetup(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "server-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	dataPath := filepath.Join(tempDir, "data")
	err = os.MkdirAll(dataPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	store, err := storage.NewBadgerStore(dataPath)
	if err != nil {
		t.Fatalf("Failed to initialize storage: %v", err)
	}
	defer store.Close()

	noteStore := storage.NewNoteStore(store)
	apiHandler := api.NewAPI(noteStore)

	r := chi.NewRouter()
	apiHandler.RegisterRoutes(r)

	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "GET /notes",
			method:         "GET",
			path:           "/notes",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "GET /note/nonexistent",
			method:         "GET",
			path:           "/note/nonexistent",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "POST /publish without body",
			method:         "POST",
			path:           "/publish",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "DELETE /note/nonexistent",
			method:         "DELETE",
			path:           "/note/nonexistent",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, w.Code)
			}
		})
	}
}

func TestCORSMiddleware(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cors-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	dataPath := filepath.Join(tempDir, "data")
	err = os.MkdirAll(dataPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	store, err := storage.NewBadgerStore(dataPath)
	if err != nil {
		t.Fatalf("Failed to initialize storage: %v", err)
	}
	defer store.Close()

	noteStore := storage.NewNoteStore(store)
	apiHandler := api.NewAPI(noteStore)

	r := chi.NewRouter()

	apiHandler.RegisterRoutes(r)

	req := httptest.NewRequest("GET", "/notes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
