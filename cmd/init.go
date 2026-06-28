package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/DycandX/flowkit/internal/config"
	"github.com/DycandX/flowkit/internal/detector"
	"github.com/DycandX/flowkit/internal/generator"
	"github.com/DycandX/flowkit/internal/prompts"
	"github.com/spf13/cobra"
)

var (
	force         bool
	configPath    string
	projectName   string
	mainBranch    string
	language      string
	workflowStyle string
	ciEnabled     bool
	hooksEnabled  bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffold workflow files for your project",
	Long: `Scaffold WORKFLOW.md, CI pipeline, and git hooks for your project.

Run without flags for interactive mode.
Pass --project-name for non-interactive.
Pass --config to re-read a saved flowkit.json.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkPrerequisites(); err != nil {
			return err
		}

		stack := string(detector.Detect())
		fmt.Printf("→ Detected stack: %s\n\n", stack)

		var cfg *config.Config

		if configPath != "" {
			var err error
			cfg, err = loadConfig(configPath, stack)
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}
		} else if projectName != "" {
			cfg = &config.Config{
				ProjectName:   projectName,
				MainBranch:    mainBranch,
				Language:      language,
				WorkflowStyle: workflowStyle,
				Stack:         stack,
				Features: config.FeaturesConfig{
					CI:          ciEnabled,
					PRCheck:     ciEnabled,
					CommitHooks: hooksEnabled,
					WorkflowDoc: true,
				},
				Commands: prompts.DefaultCommands(stack),
			}
		} else {
			var err error
			cfg, err = prompts.RunInteractive(stack)
			if err != nil {
				return err
			}
		}

		cfg.ProjectName = sanitizeName(cfg.ProjectName)

		fmt.Println("\nGenerating workflow files...")
		if err := generator.GenerateAll(cfg, force); err != nil {
			return fmt.Errorf("generation failed: %w", err)
		}

		fmt.Println("\n✅ Done! Next steps:")
		fmt.Println("  1. Review generated files")
		fmt.Println("  2. git add .")
		fmt.Println("  3. git commit -m \"chore: add workflow scaffold\"")
		fmt.Println("")
		isJS := cfg.Stack == "next" || cfg.Stack == "react" || cfg.Stack == "vue" || cfg.Stack == "nuxt" || cfg.Stack == "node"
		if cfg.Features.CommitHooks {
			if isJS {
				fmt.Println("  To activate git hooks:")
				fmt.Println("    npm install         (activates Husky)")
			} else {
				fmt.Println("  To activate git hooks:")
				fmt.Println("    git config core.hooksPath .githooks")
			}
		}
		return nil
	},
}

func checkPrerequisites() error {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("⚠  Warning: not a git repository. Run git init first.")
		fmt.Println("   Git hooks require a git repo.")
		fmt.Println()
	}
	return nil
}

func loadConfig(path, stack string) (*config.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	var cfg config.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	cfg.Stack = stack
	return &cfg, nil
}

func sanitizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "/", "-")
	name = strings.ReplaceAll(name, "\\", "-")
	name = strings.ReplaceAll(name, ":", "-")
	name = strings.ReplaceAll(name, "\"", "")
	name = strings.ReplaceAll(name, "'", "")
	return name
}

func init() {
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing files")
	initCmd.Flags().StringVar(&configPath, "config", "", "Path to flowkit.json (re-read saved config)")
	initCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "Project name (enables non-interactive mode)")
	initCmd.Flags().StringVarP(&mainBranch, "main-branch", "b", "master", "Main branch name")
	initCmd.Flags().StringVarP(&language, "language", "l", "en", "Language (en or id)")
	initCmd.Flags().StringVarP(&workflowStyle, "workflow-style", "w", "gitflow", "Workflow style (gitflow, github-flow, trunk-based)")
	initCmd.Flags().BoolVar(&ciEnabled, "ci", true, "Generate CI pipeline")
	initCmd.Flags().BoolVar(&hooksEnabled, "hooks", true, "Generate git hooks")
	rootCmd.AddCommand(initCmd)
}
