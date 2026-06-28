package config

type Config struct {
	ProjectName   string         `json:"project_name"`
	MainBranch    string         `json:"main_branch"`
	Language      string         `json:"language"`
	WorkflowStyle string         `json:"workflow_style"`
	Stack         string         `json:"stack"`
	Features      FeaturesConfig `json:"features"`
	Commands      CommandsConfig `json:"commands"`
}

type FeaturesConfig struct {
	CI          bool `json:"ci"`
	PRCheck     bool `json:"pr_check"`
	CommitHooks bool `json:"commit_hooks"`
	WorkflowDoc bool `json:"workflow_doc"`
}

type CommandsConfig struct {
	Install string `json:"install"`
	Dev     string `json:"dev"`
	Build   string `json:"build"`
	Lint    string `json:"lint"`
}
