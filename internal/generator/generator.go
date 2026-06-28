package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/DycandX/flowkit/internal/config"
)

type TemplateData struct {
	ProjectName   string
	MainBranch    string
	Stack         string
	InstallCmd    string
	DevCmd        string
	BuildCmd      string
	LintCmd       string
	WorkflowStyle string
}

var files = []struct {
	Template string
	Output   string
	SkipMsg  string
}{
	{
		Template: "templates/workflows/gitflow/WORKFLOW.md",
		Output:   "WORKFLOW.md",
		SkipMsg:  "WORKFLOW.md already exists, skipping",
	},
	{
		Template: "templates/ci/next.yml",
		Output:   ".github/workflows/ci.yml",
		SkipMsg:  ".github/workflows/ci.yml already exists, skipping",
	},
	{
		Template: "templates/hooks/pre-commit.sh",
		Output:   ".husky/pre-commit",
		SkipMsg:  ".husky/pre-commit already exists, skipping",
	},
}

func GenerateAll(cfg *config.Config, force bool) error {
	for _, f := range files {
		if err := generateFile(f.Template, f.Output, cfg, force, f.SkipMsg); err != nil {
			return err
		}
	}
	return nil
}

func generateFile(tmplPath, outputPath string, cfg *config.Config, force bool, skipMsg string) error {
	if !force {
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Printf("⏭  %s\n", skipMsg)
			return nil
		}
	}

	tmplContent, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("read template %s: %w", tmplPath, err)
	}

	data := TemplateData{
		ProjectName:   cfg.ProjectName,
		MainBranch:    cfg.MainBranch,
		Stack:         cfg.Stack,
		InstallCmd:    cfg.Commands.Install,
		DevCmd:        cfg.Commands.Dev,
		BuildCmd:      cfg.Commands.Build,
		LintCmd:       cfg.Commands.Lint,
		WorkflowStyle: cfg.WorkflowStyle,
	}

	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("parse template %s: %w", tmplPath, err)
	}

	dir := filepath.Dir(outputPath)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("create dir %s: %w", dir, err)
		}
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file %s: %w", outputPath, err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, data); err != nil {
		return fmt.Errorf("execute template %s: %w", tmplPath, err)
	}

	fmt.Printf("✓ %s\n", outputPath)
	return nil
}
