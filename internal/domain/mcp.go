package domain

type MCPServer struct {
	Capability
	Command string
	Args    []string
	Env     map[string]string
	MCPType string
	Url     string
}

type MCPServerParams struct {
	Name    string
	Scope   CapabilityScope
	Command string
	Args    []string
	Env     map[string]string
	MCPType string
	Url     string
}

func NewMCPServer(params MCPServerParams) *MCPServer {
	return &MCPServer{
		Capability: Capability{
			Name:  params.Name,
			Type:  TypeMCP,
			Scope: params.Scope,
		},
		Command: params.Command,
		Args:    params.Args,
		Env:     params.Env,
		MCPType: params.MCPType,
		Url:     params.Url,
	}
}
