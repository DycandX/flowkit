package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flowkit",
	Short: "Scaffold project workflow in one command",
	Long: `flowkit generates WORKFLOW.md, CI pipelines, and git hooks
for your project. Auto-detects stack and supports multiple workflow styles.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
