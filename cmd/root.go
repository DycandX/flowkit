package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev"

const banner = `
  ╔══════════════════════════════════════╗
  ║            flowkit                   ║
  ║    Workflow scaffold in 1 command     ║
  ╚══════════════════════════════════════╝
`

var rootCmd = &cobra.Command{
	Use:   "flowkit",
	Short: "Scaffold project workflow in one command",
	Long: banner + `
flowkit generates WORKFLOW.md, CI pipelines, and git hooks
for your project. Auto-detects stack and supports multiple workflow styles.`,
	SilenceUsage: true,
}

func Execute() {
	rootCmd.Version = Version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
