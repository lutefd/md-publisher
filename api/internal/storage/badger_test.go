package storage

import (
	"os"
	"testing"
)

func TestBadgerStore(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "badger-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	store, err := NewBadgerStore(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerStore: %v", err)
	}
	defer store.Close()

	key := "test-key"
	value := []byte("test-value")

	err = store.Set(key, value)
	if err != nil {
		t.Fatalf("Failed to set value: %v", err)
	}

	retrievedValue, err := store.Get(key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %q, got %q", value, retrievedValue)
	}
	keys, err := store.ListKeys()
	if err != nil {
		t.Fatalf("Failed to list keys: %v", err)
	}
	if len(keys) != 1 || keys[0] != key {
		t.Errorf("Expected keys [%s], got %v", key, keys)
	}

	err = store.Delete(key)
	if err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
	_, err = store.Get(key)
	if err == nil {
		t.Errorf("Expected error when getting deleted key")
	}
}

func TestNoteStore(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "notestore-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	store, err := NewBadgerStore(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerStore: %v", err)
	}
	defer store.Close()

	noteStore := NewNoteStore(store)

	note := Note{
		ID:      "test-note",
		Content: "Test content",
		Metadata: map[string]interface{}{
			"title": "Test Note",
			"tags":  []string{"test", "example"},
		},
	}

	err = noteStore.SaveNote(note)
	if err != nil {
		t.Fatalf("Failed to save note: %v", err)
	}

	retrievedNote, err := noteStore.GetNote(note.ID)
	if err != nil {
		t.Fatalf("Failed to get note: %v", err)
	}
	if retrievedNote.ID != note.ID {
		t.Errorf("Expected note ID %q, got %q", note.ID, retrievedNote.ID)
	}
	if retrievedNote.Content != note.Content {
		t.Errorf("Expected note content %q, got %q", note.Content, retrievedNote.Content)
	}
	if retrievedNote.Metadata["title"] != note.Metadata["title"] {
		t.Errorf("Expected note title %q, got %q", note.Metadata["title"], retrievedNote.Metadata["title"])
	}

	notes, err := noteStore.ListNotes()
	if err != nil {
		t.Fatalf("Failed to list notes: %v", err)
	}
	if len(notes) != 1 || notes[0].ID != note.ID {
		t.Errorf("Expected notes with ID [%s], got %v", note.ID, notes)
	}

	err = noteStore.DeleteNote(note.ID)
	if err != nil {
		t.Fatalf("Failed to delete note: %v", err)
	}
	_, err = noteStore.GetNote(note.ID)
	if err == nil {
		t.Errorf("Expected error when getting deleted note")
	}
}

func TestNoteStoreWithFrontmatter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "notestore-frontmatter-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	store, err := NewBadgerStore(tempDir)
	if err != nil {
		t.Fatalf("Failed to create BadgerStore: %v", err)
	}
	defer store.Close()

	noteStore := NewNoteStore(store)

	noteWithFrontmatter := Note{
		ID: "frontmatter-note",
		Content: `---
title: Frontmatter Title
description: This is a description from frontmatter
tags:
  - tag1
  - tag2
author: Test Author
---
# Actual Content

This is the content of the note.`,
		Metadata: map[string]interface{}{
			"existing": "metadata",
		},
	}

	err = noteStore.SaveNote(noteWithFrontmatter)
	if err != nil {
		t.Fatalf("Failed to save note with frontmatter: %v", err)
	}

	retrievedNote, err := noteStore.GetNote(noteWithFrontmatter.ID)
	if err != nil {
		t.Fatalf("Failed to get note with frontmatter: %v", err)
	}
	if retrievedNote.Metadata["title"] != "Frontmatter Title" {
		t.Errorf("Expected title from frontmatter, got %v", retrievedNote.Metadata["title"])
	}
	if retrievedNote.Metadata["description"] != "This is a description from frontmatter" {
		t.Errorf("Expected description from frontmatter, got %v", retrievedNote.Metadata["description"])
	}
	if retrievedNote.Metadata["author"] != "Test Author" {
		t.Errorf("Expected author from frontmatter, got %v", retrievedNote.Metadata["author"])
	}

	if retrievedNote.Metadata["existing"] != "metadata" {
		t.Errorf("Expected existing metadata to be preserved, got %v", retrievedNote.Metadata["existing"])
	}
	expectedContent := "# Actual Content\n\nThis is the content of the note."
	if retrievedNote.Content != expectedContent {
		t.Errorf("Expected content without frontmatter, got %q", retrievedNote.Content)
	}
}
