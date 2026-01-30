package loaders

import (
	"bytes"

	"claudectl/internal/domain"
)

type Loader[T any] interface {
	Load(scope domain.CapabilityScope) ([]T, error)
}

type MarkdownMetadata struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func hasYAMLFrontmatter(content []byte) bool {
	return bytes.HasPrefix(content, []byte("---\n")) || bytes.HasPrefix(content, []byte("---\r\n"))
}

func extractYAMLFrontmatterAndContent(content []byte) ([]byte, []byte) {
	if !hasYAMLFrontmatter(content) {
		return nil, content
	}

	lines := bytes.Split(content, []byte("\n"))
	endIdx := -1
	for i := 1; i < len(lines); i++ {
		line := bytes.TrimSpace(lines[i])
		if bytes.Equal(line, []byte("---")) {
			endIdx = i
			break
		}
	}

	if endIdx == -1 {
		return nil, content
	}

	frontmatter := bytes.Join(lines[1:endIdx], []byte("\n"))
	body := bytes.Join(lines[endIdx+1:], []byte("\n"))

	return frontmatter, body
}

func parseMarkdownWithFrontmatter(content []byte) (*MarkdownMetadata, []byte, error) {
	frontmatter, contentBody := extractYAMLFrontmatterAndContent(content)
	if frontmatter == nil {
		return nil, contentBody, nil
	}

	var metadata MarkdownMetadata

	metadata = parseKeyValue(frontmatter)

	return &metadata, contentBody, nil
}

func parseKeyValue(data []byte) MarkdownMetadata {
	var metadata MarkdownMetadata
	lines := bytes.Split(data, []byte("\n"))

	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		// Find first colon
		colonIdx := bytes.IndexByte(line, ':')
		if colonIdx == -1 {
			continue
		}

		key := string(bytes.TrimSpace(line[:colonIdx]))
		value := string(bytes.TrimSpace(line[colonIdx+1:]))

		switch key {
		case "name":
			metadata.Name = value
		case "description":
			metadata.Description = value
		}
	}

	return metadata
}
