package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev"

const banner = `
 ╔══════════════════════════════════════╗
 ║            flowkit                    ║
 ║    Workflow scaffold in 1 command     ║
 ╚══════════════════════════════════════╝
`

var rootCmd = &cobra.Command{
	Use:   "flowkit",
	Short: "Scaffold project workflow in one command",
	Long: fmt.Sprintf(`%s
flowkit generates WORKFLOW.md, CI pipelines, and git hooks
for your project. Auto-detects stack and supports multiple workflow styles.`, banner),
	SilenceUsage: true,
	Version:      Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
