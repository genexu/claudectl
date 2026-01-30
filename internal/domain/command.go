package domain

type Command struct {
	Capability
	FilePath string
	Content  string
}

type CommandParams struct {
	Name        string
	Description string
	Scope       CapabilityScope
	FilePath    string
	Content     string
}

func NewCommand(params CommandParams) *Command {
	return &Command{
		Capability: Capability{
			Name:        params.Name,
			Description: params.Description,
			Type:        TypeCommand,
			Scope:       params.Scope,
		},
		FilePath: params.FilePath,
		Content:  params.Content,
	}
}
