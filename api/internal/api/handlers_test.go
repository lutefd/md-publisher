package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lutefd/md-publsiher/api/internal/storage"
)

type MockNoteStore struct {
	notes map[string]storage.Note
}

func NewMockNoteStore() *MockNoteStore {
	return &MockNoteStore{
		notes: make(map[string]storage.Note),
	}
}

func (m *MockNoteStore) SaveNote(note storage.Note) error {
	storage.ExtractFrontmatter(&note)
	m.notes[note.ID] = note
	return nil
}

func (m *MockNoteStore) GetNote(id string) (storage.Note, error) {
	note, exists := m.notes[id]
	if !exists {
		return storage.Note{}, errors.New("note not found")
	}
	return note, nil
}

func (m *MockNoteStore) DeleteNote(id string) error {
	if _, exists := m.notes[id]; !exists {
		return errors.New("note not found")
	}
	delete(m.notes, id)
	return nil
}

func (m *MockNoteStore) ListNotes() ([]storage.Note, error) {
	notes := make([]storage.Note, 0, len(m.notes))
	for _, note := range m.notes {
		notes = append(notes, note)
	}
	return notes, nil
}

func TestPublishNote(t *testing.T) {
	mockStore := NewMockNoteStore()

	api := NewAPI(mockStore)
	note := storage.Note{
		ID:      "test-note",
		Content: "Test content",
		Metadata: map[string]interface{}{
			"title": "Test Note",
		},
	}

	noteJSON, err := json.Marshal(note)
	if err != nil {
		t.Fatalf("Failed to marshal note: %v", err)
	}

	req := httptest.NewRequest("POST", "/publish", bytes.NewBuffer(noteJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	api.PublishNote(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	savedNote, err := mockStore.GetNote(note.ID)
	if err != nil {
		t.Fatalf("Failed to get saved note: %v", err)
	}
	if savedNote.ID != note.ID {
		t.Errorf("Expected note ID %q, got %q", note.ID, savedNote.ID)
	}
}

func TestPublishNoteWithFrontmatter(t *testing.T) {
	// Create a mock note store
	mockStore := NewMockNoteStore()

	// Create an API instance with the mock store
	api := NewAPI(mockStore)

	// Create a test note with frontmatter
	note := storage.Note{
		ID: "frontmatter-note",
		Content: `---
title: Frontmatter Title
tags:
  - tag1
  - tag2
---
# Content`,
		Metadata: map[string]interface{}{
			"existing": "metadata",
		},
	}

	// Convert note to JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		t.Fatalf("Failed to marshal note: %v", err)
	}

	// Create a request
	req := httptest.NewRequest("POST", "/publish", bytes.NewBuffer(noteJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	api.PublishNote(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify the note was saved with frontmatter extracted
	savedNote, err := mockStore.GetNote(note.ID)
	if err != nil {
		t.Fatalf("Failed to get saved note: %v", err)
	}

	// Check that frontmatter was extracted
	if savedNote.Metadata["title"] != "Frontmatter Title" {
		t.Errorf("Expected title from frontmatter, got %v", savedNote.Metadata["title"])
	}

	// Check that content no longer has frontmatter
	expectedContent := "# Content"
	if savedNote.Content != expectedContent {
		t.Errorf("Expected content without frontmatter, got %q", savedNote.Content)
	}
}

func TestGetNote(t *testing.T) {
	// Create a mock note store
	mockStore := NewMockNoteStore()

	// Create an API instance with the mock store
	api := NewAPI(mockStore)

	// Add a test note to the store
	testNote := storage.Note{
		ID:      "test-note",
		Content: "Test content",
		Metadata: map[string]interface{}{
			"title": "Test Note",
		},
	}
	mockStore.SaveNote(testNote)

	// Create a new router
	r := chi.NewRouter()
	r.Get("/note/{id}", api.GetNote)

	// Create a request
	req := httptest.NewRequest("GET", "/note/test-note", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response
	var responseNote storage.Note
	err := json.Unmarshal(w.Body.Bytes(), &responseNote)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify the note
	if responseNote.ID != testNote.ID {
		t.Errorf("Expected note ID %q, got %q", testNote.ID, responseNote.ID)
	}
	if responseNote.Content != testNote.Content {
		t.Errorf("Expected note content %q, got %q", testNote.Content, responseNote.Content)
	}
}

func TestListNotes(t *testing.T) {
	// Create a mock note store
	mockStore := NewMockNoteStore()

	// Create an API instance with the mock store
	api := NewAPI(mockStore)

	// Add test notes to the store
	testNotes := []storage.Note{
		{
			ID:      "note1",
			Content: "Content 1",
			Metadata: map[string]interface{}{
				"title": "Note 1",
			},
		},
		{
			ID:      "note2",
			Content: "Content 2",
			Metadata: map[string]interface{}{
				"title": "Note 2",
			},
		},
	}

	for _, note := range testNotes {
		mockStore.SaveNote(note)
	}

	// Create a request
	req := httptest.NewRequest("GET", "/notes", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	api.ListNotes(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response
	var responseNotes []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseNotes)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify the notes
	if len(responseNotes) != len(testNotes) {
		t.Errorf("Expected %d notes, got %d", len(testNotes), len(responseNotes))
	}
}

func TestUnpublishNote(t *testing.T) {
	// Create a mock note store
	mockStore := NewMockNoteStore()

	// Create an API instance with the mock store
	api := NewAPI(mockStore)

	// Add a test note to the store
	testNote := storage.Note{
		ID:      "test-note",
		Content: "Test content",
	}
	mockStore.SaveNote(testNote)

	// Create a new router
	r := chi.NewRouter()
	r.Delete("/note/{id}", api.UnpublishNote)

	// Create a request
	req := httptest.NewRequest("DELETE", "/note/test-note", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify the note was deleted
	_, err := mockStore.GetNote(testNote.ID)
	if err == nil {
		t.Errorf("Expected note to be deleted")
	}
}
