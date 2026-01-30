package loaders

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"claudectl/internal/domain"
	"claudectl/internal/utils"
)

type AgentLoader struct {
	logger *slog.Logger
}

func NewAgentLoader(logger *utils.Logger) Loader[domain.Agent] {
	logger.Debug("initializing agent loader")
	return &AgentLoader{logger: logger.Logger}
}

func (a *AgentLoader) Load(scope domain.CapabilityScope) ([]domain.Agent, error) {
	basePath, err := utils.GetScopeBaseDir(scope)
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(basePath, "agents")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		a.logger.Debug("directory not found", "path", dir)
		return []domain.Agent{}, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		a.logger.Error("failed to read directory", "path", dir, "error", err)
		return nil, err
	}

	var agents []domain.Agent
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())

		content, err := os.ReadFile(filePath)
		if err != nil {
			a.logger.Warn("failed to read file", "path", filePath, "error", err)
			continue
		}

		metadata, body, err := parseMarkdownWithFrontmatter(content)
		if err != nil {
			a.logger.Warn("failed to parse frontmatter", "path", filePath, "error", err)
		}

		name := strings.TrimSuffix(entry.Name(), ".md")
		if metadata != nil && metadata.Name != "" {
			name = metadata.Name
		}

		description := ""
		if metadata != nil {
			description = metadata.Description
		}

		agent := domain.NewAgent(domain.AgentParams{
			Name:        name,
			Description: description,
			FilePath:    filePath,
			Content:     string(body),
			Scope:       scope,
		})
		agents = append(agents, *agent)
		a.logger.Debug("discovered agent", "name", name, "scope", scope)
	}

	a.logger.Info("discovered agents", "count", len(agents), "path", dir)
	return agents, nil
}
