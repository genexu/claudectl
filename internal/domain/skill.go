package domain

type Skill struct {
	Capability
	FilePath string
	Content  string
}

type SkillParams struct {
	Name        string
	Description string
	Scope       CapabilityScope
	FilePath    string
	Content     string
}

func NewSkill(params SkillParams) *Skill {
	return &Skill{
		Capability: Capability{
			Name:        params.Name,
			Description: params.Description,
			Type:        TypeSkill,
			Scope:       params.Scope,
		},
		FilePath: params.FilePath,
		Content:  params.Content,
	}
}
