package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/lutefd/md-publsiher/api/internal/storage"
)

type NoteStorer interface {
	SaveNote(note storage.Note) error
	GetNote(id string) (storage.Note, error)
	DeleteNote(id string) error
	ListNotes() ([]storage.Note, error)
}

type API struct {
	noteStore NoteStorer
}

func NewAPI(noteStore NoteStorer) *API {
	return &API{
		noteStore: noteStore,
	}
}

func (api *API) RegisterRoutes(r chi.Router) {
	r.Post("/publish", api.PublishNote)
	r.Delete("/note/{id}", api.UnpublishNote)
	r.Get("/notes", api.ListNotes)
	r.Get("/note/{id}", api.GetNote)
}
func (api *API) PublishNote(w http.ResponseWriter, r *http.Request) {
	var note storage.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if note.ID == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}

	if err := api.noteStore.SaveNote(note); err != nil {
		http.Error(w, "Failed to store note", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"status": "Note published successfully"})
}
func (api *API) UnpublishNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}

	if err := api.noteStore.DeleteNote(id); err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"status": "Note unpublished successfully"})
}

func (api *API) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := api.noteStore.ListNotes()
	if err != nil {
		http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
		return
	}

	response := make([]map[string]interface{}, 0, len(notes))
	for _, note := range notes {
		noteResponse := map[string]interface{}{
			"id":       note.ID,
			"content":  note.Content,
			"metadata": note.Metadata,
		}

		if _, exists := note.Metadata["updated"]; !exists {
			if noteResponse["metadata"] == nil {
				noteResponse["metadata"] = make(map[string]interface{})
			}
			noteResponse["metadata"].(map[string]interface{})["updated"] = time.Now().Format(time.RFC3339)
		}

		response = append(response, noteResponse)
	}

	render.JSON(w, r, response)
}

func (api *API) GetNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}

	note, err := api.noteStore.GetNote(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"id":       note.ID,
		"content":  note.Content,
		"metadata": note.Metadata,
	}

	if _, exists := note.Metadata["updated"]; !exists {
		if response["metadata"] == nil {
			response["metadata"] = make(map[string]interface{})
		}
		response["metadata"].(map[string]interface{})["updated"] = time.Now().Format(time.RFC3339)
	}

	render.JSON(w, r, response)
}
