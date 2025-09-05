package storage

import (
	"reflect"
	"testing"
)

func TestParseFrontmatter(t *testing.T) {
	tests := []struct {
		name            string
		content         string
		wantFrontmatter map[string]interface{}
		wantContent     string
	}{
		{
			name: "Valid frontmatter",
			content: `---
autor: "André Almeida"
iniciativa: "[SPM] Capability Eshops"
tags:
  - Frontend
  - Backend
cssclasses:
  - center-titles
  - center-images
---
# Migração de EshopsItem para Polycards em Super Shops
***
## Resumo

Content here...`,
			wantFrontmatter: map[string]interface{}{
				"autor":      "André Almeida",
				"iniciativa": "[SPM] Capability Eshops",
				"tags":       []interface{}{"Frontend", "Backend"},
				"cssclasses": []interface{}{"center-titles", "center-images"},
			},
			wantContent: "# Migração de EshopsItem para Polycards em Super Shops\n***\n## Resumo\n\nContent here...",
		},
		{
			name:            "No frontmatter",
			content:         "# Just a regular markdown\n\nWith no frontmatter",
			wantFrontmatter: nil,
			wantContent:     "# Just a regular markdown\n\nWith no frontmatter",
		},
		{
			name: "Invalid frontmatter",
			content: `---
invalid: yaml: [
---
# Content with invalid frontmatter`,
			wantFrontmatter: nil,
			wantContent:     "---\ninvalid: yaml: [\n---\n# Content with invalid frontmatter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFrontmatter, gotContent := ParseFrontmatter(tt.content)

			if !reflect.DeepEqual(gotFrontmatter, tt.wantFrontmatter) {
				t.Errorf("ParseFrontmatter() frontmatter = %v, want %v", gotFrontmatter, tt.wantFrontmatter)
			}
			if gotContent != tt.wantContent {
				t.Errorf("ParseFrontmatter() content = %v, want %v", gotContent, tt.wantContent)
			}
		})
	}
}

func TestExtractFrontmatter(t *testing.T) {
	tests := []struct {
		name         string
		note         Note
		wantMetadata map[string]interface{}
		wantContent  string
	}{
		{
			name: "Note with frontmatter",
			note: Note{
				ID: "test-note",
				Content: `---
title: Test Note
tags:
  - test
  - example
---
# Test Content`,
				Metadata: map[string]interface{}{
					"existing": "metadata",
				},
			},
			wantMetadata: map[string]interface{}{
				"existing": "metadata",
				"title":    "Test Note",
				"tags":     []interface{}{"test", "example"},
			},
			wantContent: "# Test Content",
		},
		{
			name: "Note without frontmatter",
			note: Note{
				ID:       "test-note",
				Content:  "# Just content",
				Metadata: map[string]interface{}{"existing": "metadata"},
			},
			wantMetadata: map[string]interface{}{"existing": "metadata"},
			wantContent:  "# Just content",
		},
		{
			name: "Note with nil metadata",
			note: Note{
				ID: "test-note",
				Content: `---
title: Test Note
---
# Content`,
				Metadata: nil,
			},
			wantMetadata: map[string]interface{}{
				"title": "Test Note",
			},
			wantContent: "# Content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note := tt.note
			ExtractFrontmatter(&note)

			if !reflect.DeepEqual(note.Metadata, tt.wantMetadata) {
				t.Errorf("ExtractFrontmatter() metadata = %v, want %v", note.Metadata, tt.wantMetadata)
			}
			if note.Content != tt.wantContent {
				t.Errorf("ExtractFrontmatter() content = %v, want %v", note.Content, tt.wantContent)
			}
		})
	}
}
