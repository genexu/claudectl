package domain

type PluginAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	URL   string `json:"url,omitempty"`
}

type Plugin struct {
	Capability
	Version    string       `json:"version"`
	Author     PluginAuthor `json:"author"`
	Homepage   string       `json:"homepage,omitempty"`
	Repository string       `json:"repository,omitempty"`
	License    string       `json:"license,omitempty"`
	Keywords   []string     `json:"keywords,omitempty"`
	MCPServers []MCPServer  `json:"mcp_servers,omitempty"`
	Commands   []Command    `json:"commands,omitempty"`
	Skills     []Skill      `json:"skills,omitempty"`
	Agents     []Agent      `json:"agents,omitempty"`
	Path       string       `json:"path,omitempty"`
}

func (p *Plugin) CapabilityCount() int {
	return len(p.MCPServers) + len(p.Commands) +
		len(p.Skills) + len(p.Agents)
}

func (p *Plugin) GetType() CapabilityType {
	return TypePlugin
}

func (p *Plugin) GetScope() CapabilityScope {
	return p.Scope
}

func (p *Plugin) GetName() string {
	return p.Name
}
