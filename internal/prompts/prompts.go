package prompts

import (
	"github.com/charmbracelet/huh"
	"github.com/DycandX/flowkit/internal/config"
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
		Commands: DefaultCommands(stack),
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name").
				Value(&cfg.ProjectName).
				Validate(func(s string) error {
					if s == "" {
						return nil // allowed to be empty here, validated later
					}
					return nil
				}),

			huh.NewInput().
				Title("Main branch").
				Value(&cfg.MainBranch).
				Description("Default: master"),

			huh.NewSelect[string]().
				Title("Language").
				Options(
					huh.NewOption("English", "en"),
					huh.NewOption("Bahasa Indonesia", "id"),
				).
				Value(&cfg.Language),

			huh.NewSelect[string]().
				Title("Workflow style").
				Options(
					huh.NewOption("GitFlow", "gitflow"),
					huh.NewOption("GitHub Flow", "github-flow"),
					huh.NewOption("Trunk-Based", "trunk-based"),
				).
				Value(&cfg.WorkflowStyle),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func DefaultCommands(stack string) config.CommandsConfig {
	switch stack {
	case "go":
		return config.CommandsConfig{
			Install: "go mod download",
			Dev:     "go run .",
			Build:   "go build .",
			Lint:    "go vet ./...",
		}
	case "rust":
		return config.CommandsConfig{
			Install: "cargo build",
			Dev:     "cargo run",
			Build:   "cargo build --release",
			Lint:    "cargo clippy",
		}
	case "python":
		return config.CommandsConfig{
			Install: "pip install -r requirements.txt",
			Dev:     "python .",
			Build:   "python -m build",
			Lint:    "ruff check .",
		}
	case "laravel":
		return config.CommandsConfig{
			Install: "composer install",
			Dev:     "php artisan serve",
			Build:   "npm run build && php artisan optimize",
			Lint:    "php artisan lint",
		}
	default:
		return config.CommandsConfig{
			Install: "npm install",
			Dev:     "npm run dev",
			Build:   "npm run build",
			Lint:    "npm run lint",
		}
	}
}
