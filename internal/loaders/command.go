package loaders

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"claudectl/internal/domain"
	"claudectl/internal/utils"
)

type CommandLoader struct {
	logger *slog.Logger
}

func NewCommandLoader(logger *utils.Logger) Loader[domain.Command] {
	logger.Debug("initializing command loader")
	return &CommandLoader{logger: logger.Logger}
}

func (c *CommandLoader) Load(scope domain.CapabilityScope) ([]domain.Command, error) {
	basePath, err := utils.GetScopeBaseDir(scope)
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(basePath, "commands")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		c.logger.Debug("directory not found", "path", dir)
		return []domain.Command{}, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		c.logger.Error("failed to read directory", "path", dir, "error", err)
		return nil, err
	}

	var commands []domain.Command
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())

		content, err := os.ReadFile(filePath)
		if err != nil {
			c.logger.Warn("failed to read file", "path", filePath, "error", err)
			continue
		}

		metadata, body, err := parseMarkdownWithFrontmatter(content)
		if err != nil {
			c.logger.Warn("failed to parse frontmatter", "path", filePath, "error", err)
		}

		name := strings.TrimSuffix(entry.Name(), ".md")
		if metadata != nil && metadata.Name != "" {
			name = metadata.Name
		}

		description := ""
		if metadata != nil {
			description = metadata.Description
		}

		command := domain.NewCommand(domain.CommandParams{
			Name:        name,
			Description: description,
			FilePath:    filePath,
			Content:     string(body),
			Scope:       scope,
		})
		commands = append(commands, *command)
		c.logger.Debug("discovered command", "name", name, "scope", scope)
	}

	c.logger.Info("discovered commands", "count", len(commands), "path", dir)
	return commands, nil
}
