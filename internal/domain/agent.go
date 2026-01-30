package domain

type Agent struct {
	Capability
	FilePath string
	Content  string
}

type AgentParams struct {
	Name        string
	Scope       CapabilityScope
	FilePath    string
	Description string
	Content     string
}

func NewAgent(params AgentParams) *Agent {
	return &Agent{
		Capability: Capability{
			Name:        params.Name,
			Description: params.Description,
			Type:        TypeAgent,
			Scope:       params.Scope,
		},
		FilePath: params.FilePath,
		Content:  params.Content,
	}
}
