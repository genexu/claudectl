package loaders

import (
	"log/slog"
	"os"
	"path/filepath"

	"claudectl/internal/domain"
	"claudectl/internal/utils"
)

type SkillLoader struct {
	logger *slog.Logger
}

func NewSkillLoader(logger *utils.Logger) Loader[domain.Skill] {
	logger.Debug("initializing skill loader")
	return &SkillLoader{logger: logger.Logger}
}

func (s *SkillLoader) Load(scope domain.CapabilityScope) ([]domain.Skill, error) {
	basePath, err := utils.GetScopeBaseDir(scope)
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(basePath, "skills")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		s.logger.Debug("directory not found", "path", dir)
		return []domain.Skill{}, nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		s.logger.Error("failed to read directory", "path", dir, "error", err)
		return nil, err
	}

	var skills []domain.Skill
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skillDir := entry.Name()
		skillFilePath := filepath.Join(dir, skillDir, "SKILL.md")

		// Check if SKILL.md exists
		if _, err := os.Stat(skillFilePath); os.IsNotExist(err) {
			continue
		}

		content, err := os.ReadFile(skillFilePath)
		if err != nil {
			s.logger.Warn("failed to read file", "path", skillFilePath, "error", err)
			continue
		}

		metadata, body, err := parseMarkdownWithFrontmatter(content)
		if err != nil {
			s.logger.Warn("failed to parse frontmatter", "path", skillFilePath, "error", err)
		}

		name := skillDir
		if metadata != nil && metadata.Name != "" {
			name = metadata.Name
		}

		description := ""
		if metadata != nil {
			description = metadata.Description
		}

		skill := domain.NewSkill(domain.SkillParams{
			Name:        name,
			Description: description,
			FilePath:    skillFilePath,
			Content:     string(body),
			Scope:       scope,
		})
		skills = append(skills, *skill)
		s.logger.Debug("discovered skill", "name", name, "scope", scope)
	}

	s.logger.Info("discovered skills", "count", len(skills), "path", dir)
	return skills, nil
}
