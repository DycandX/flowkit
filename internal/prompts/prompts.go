package prompts

import (
	"fmt"

	"github.com/DycandX/flowkit/internal/config"
	"github.com/manifoldco/promptui"
)

func RunInteractive(stack string) (*config.Config, error) {
	cfg := &config.Config{
		MainBranch:    "master",
		Language:      "en",
		WorkflowStyle: "gitflow",
		Stack:         stack,
		Features: config.FeaturesConfig{
			CI:          true,
			PRCheck:     true,
			CommitHooks: true,
			WorkflowDoc: true,
		},
		Commands: config.CommandsConfig{
			Install: "npm install",
			Dev:     "npm run dev",
			Build:   "npm run build",
			Lint:    "npm run lint",
		},
	}

	var err error

	cfg.ProjectName, err = promptString("Project name", "")
	if err != nil {
		return nil, err
	}

	cfg.MainBranch, err = promptString("Main branch", cfg.MainBranch)
	if err != nil {
		return nil, err
	}

	lang, err := promptSelect("Language", []string{"en", "id"})
	if err != nil {
		return nil, err
	}
	cfg.Language = lang

	style, err := promptSelect("Workflow style", []string{"gitflow", "github-flow", "trunk-based"})
	if err != nil {
		return nil, err
	}
	cfg.WorkflowStyle = style

	fmt.Println()
	return cfg, nil
}

func promptString(label, def string) (string, error) {
	p := promptui.Prompt{
		Label:   label,
		Default: def,
	}
	return p.Run()
}

func promptSelect(label string, items []string) (string, error) {
	p := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := p.Run()
	return result, err
}
