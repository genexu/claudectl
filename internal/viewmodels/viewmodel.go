package viewmodels

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"

	"claudectl/internal/domain"
)

// CapabilityViewModel defines the common interface for all capability view models
// It combines Bubble Tea list.Item interface with domain model accessors
type CapabilityViewModel interface {
	list.Item

	GetName() string
	GetDescription() string
	GetScope() domain.CapabilityScope
	GetType() domain.CapabilityType
	GetFilePath() string
	GetContent() string

	RenderDetails() []string
}

// ToDomainViewModel converts a domain model to a view model
// Note: Only handles value types since loaders return slices of values
func ToDomainViewModel(cap any) (CapabilityViewModel, error) {
	switch v := cap.(type) {
	case domain.MCPServer:
		return NewMCPServerViewModel(&v), nil
	case domain.Command:
		return NewCommandViewModel(&v), nil
	case domain.Skill:
		return NewSkillViewModel(&v), nil
	case domain.Agent:
		return NewAgentViewModel(&v), nil
	case domain.Plugin:
		return NewPluginViewModel(&v), nil
	default:
		return nil, fmt.Errorf("unsupported capability type: %T", cap)
	}
}
