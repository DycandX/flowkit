package cmd

import (
	"fmt"

	"github.com/DycandX/flowkit/internal/detector"
	"github.com/DycandX/flowkit/internal/generator"
	"github.com/DycandX/flowkit/internal/prompts"
	"github.com/spf13/cobra"
)

var force bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffold workflow files for your project",
	Long: `Interactively scaffold WORKFLOW.md, CI pipeline, and git hooks
for your project. Run from your project root directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stack := string(detector.Detect())
		fmt.Printf("→ Detected stack: %s\n\n", stack)

		cfg, err := prompts.RunInteractive(stack)
		if err != nil {
			return err
		}

		fmt.Println("\nGenerating workflow files...")
		if err := generator.GenerateAll(cfg, force); err != nil {
			return fmt.Errorf("generation failed: %w", err)
		}

		fmt.Println("\n✅ Done! Next steps:")
		fmt.Println("  1. Review generated files")
		fmt.Println("  2. git add .")
		fmt.Println("  3. git commit -m \"chore: add workflow scaffold\"")
		return nil
	},
}

func init() {
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing files")
	rootCmd.AddCommand(initCmd)
}
