package domain

type CapabilityScope string
type CapabilityType string

const (
	ScopeUser    CapabilityScope = "user"
	ScopeProject CapabilityScope = "project"
)

const (
	TypeMCP     CapabilityType = "mcp"
	TypeCommand CapabilityType = "command"
	TypeSkill   CapabilityType = "skill"
	TypeAgent   CapabilityType = "agent"
	TypePlugin  CapabilityType = "plugin"
)

type Capability struct {
	Name        string
	Description string
	Type        CapabilityType
	Scope       CapabilityScope
}
