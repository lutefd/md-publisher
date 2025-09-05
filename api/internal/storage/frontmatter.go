package storage

import (
	"bytes"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var frontmatterRegex = regexp.MustCompile(`(?s)^---\n(.*?)\n---\n(.*)$`)

func ParseFrontmatter(content string) (map[string]interface{}, string) {
	matches := frontmatterRegex.FindStringSubmatch(content)
	if len(matches) != 3 {
		return nil, content
	}

	frontmatterYAML := matches[1]
	remainingContent := matches[2]

	var frontmatter map[string]interface{}
	decoder := yaml.NewDecoder(bytes.NewBufferString(frontmatterYAML))
	err := decoder.Decode(&frontmatter)
	if err != nil {
		return nil, content
	}

	return frontmatter, strings.TrimSpace(remainingContent)
}

func ExtractFrontmatter(note *Note) {
	frontmatter, content := ParseFrontmatter(note.Content)
	if frontmatter != nil {
		if note.Metadata == nil {
			note.Metadata = make(map[string]interface{})
		}

		for key, value := range frontmatter {
			note.Metadata[key] = value
		}
		note.Content = content
	}
}
