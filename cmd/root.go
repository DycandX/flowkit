package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev"

func displayBanner() string {
	v := Version
	if v != "dev" {
		v = " " + v
	} else {
		v = ""
	}
	return fmt.Sprintf(`
  ╔══════════════════════════════════════╗
  ║            flowkit%-21s║
  ║    Workflow scaffold in 1 command     ║
  ╚══════════════════════════════════════╝
`, v)
}

var rootCmd = &cobra.Command{
	Use:   "flowkit",
	Short: "Scaffold project workflow in one command",
	Long: fmt.Sprintf(`%s
flowkit generates WORKFLOW.md, CI pipelines, and git hooks
for your project. Auto-detects stack and supports multiple workflow styles.`, displayBanner()),
	SilenceUsage: true,
	Version:      Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
