package generator

import (
	"encoding/json"
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

type fileSpec struct {
	Template string
	Output   string
	SkipMsg  string
	Cond     func(*config.Config) bool
}

func isJS(stack string) bool {
	return stack == "next" || stack == "react" || stack == "vue" || stack == "nuxt" || stack == "node"
}

func filesFor(cfg *config.Config) []fileSpec {
	ciTemplates := map[string]string{
		"next":    "templates/ci/next.yml",
		"react":   "templates/ci/react.yml",
		"vue":     "templates/ci/node.yml",
		"nuxt":    "templates/ci/next.yml",
		"node":    "templates/ci/node.yml",
		"go":      "templates/ci/go.yml",
		"rust":    "templates/ci/rust.yml",
		"python":  "templates/ci/python.yml",
		"laravel": "templates/ci/laravel.yml",
	}
	ciTmpl, ok := ciTemplates[cfg.Stack]
	if !ok {
		ciTmpl = "templates/ci/generic.yml"
	}

	workflowTmpl := fmt.Sprintf("templates/workflows/%s/WORKFLOW.md", cfg.WorkflowStyle)

	precommitTmpl := "templates/hooks/pre-commit.sh"
	precommitOut := ".husky/pre-commit"
	if !isJS(cfg.Stack) {
		precommitTmpl = "templates/hooks/pre-commit-nonjs.sh"
		precommitOut = ".githooks/pre-commit"
	}

	commitmsgOut := ".husky/commit-msg"
	commitmsgTmpl := "templates/hooks/commit-msg.sh"
	if !isJS(cfg.Stack) {
		commitmsgOut = ".githooks/commit-msg"
	}

	ff := []fileSpec{
		{
			Template: workflowTmpl,
			Output:   "WORKFLOW.md",
			SkipMsg:  "WORKFLOW.md already exists, skipping",
			Cond:     func(c *config.Config) bool { return c.Features.WorkflowDoc },
		},
		{
			Template: ciTmpl,
			Output:   ".github/workflows/ci.yml",
			SkipMsg:  ".github/workflows/ci.yml already exists, skipping",
			Cond:     func(c *config.Config) bool { return c.Features.CI },
		},
		{
			Template: "templates/ci/pr-check.yml",
			Output:   ".github/workflows/pr-check.yml",
			SkipMsg:  ".github/workflows/pr-check.yml already exists, skipping",
			Cond:     func(c *config.Config) bool { return c.Features.PRCheck },
		},
		{
			Template: precommitTmpl,
			Output:   precommitOut,
			SkipMsg:  fmt.Sprintf("%s already exists, skipping", precommitOut),
			Cond:     func(c *config.Config) bool { return c.Features.CommitHooks },
		},
		{
			Template: commitmsgTmpl,
			Output:   commitmsgOut,
			SkipMsg:  fmt.Sprintf("%s already exists, skipping", commitmsgOut),
			Cond:     func(c *config.Config) bool { return c.Features.CommitHooks },
		},
		{
			Template: "templates/commitlint.config.js",
			Output:   "commitlint.config.js",
			SkipMsg:  "commitlint.config.js already exists, skipping",
			Cond:     func(c *config.Config) bool { return c.Features.CommitHooks && isJS(cfg.Stack) },
		},
		{
			Template: "templates/lintstagedrc.json",
			Output:   ".lintstagedrc.json",
			SkipMsg:  ".lintstagedrc.json already exists, skipping",
			Cond:     func(c *config.Config) bool { return c.Features.CommitHooks && isJS(cfg.Stack) },
		},
	}

	saveConfig(cfg)
	return ff
}

func saveConfig(cfg *config.Config) {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Printf("⚠  Warning: could not marshal config: %v\n", err)
		return
	}
	if err := os.WriteFile("flowkit.json", data, 0644); err != nil {
		fmt.Printf("⚠  Warning: could not save flowkit.json: %v\n", err)
		return
	}
	fmt.Println("✓ flowkit.json (config saved)")
}

func GenerateAll(cfg *config.Config, force bool) error {
	for _, f := range filesFor(cfg) {
		if !f.Cond(cfg) {
			continue
		}
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

	tmpl, err := template.New(filepath.Base(tmplPath)).Delims("<<", ">>").Parse(string(tmplContent))
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
