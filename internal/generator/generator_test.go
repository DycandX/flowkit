package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DycandX/flowkit/internal/config"
)

func TestGenerateAll(t *testing.T) {
	cfg := &config.Config{
		ProjectName:   "test-app",
		MainBranch:    "master",
		Language:      "en",
		WorkflowStyle: "gitflow",
		Stack:         "next",
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

	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(origDir)

	if err := GenerateAll(cfg, true); err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	checkFile(t, tmpDir, "WORKFLOW.md")
	checkFile(t, tmpDir, ".github/workflows/ci.yml")
	checkFile(t, tmpDir, ".github/workflows/pr-check.yml")
	checkFile(t, tmpDir, ".husky/pre-commit")
	checkFile(t, tmpDir, ".husky/commit-msg")
	checkFile(t, tmpDir, "commitlint.config.js")
	checkFile(t, tmpDir, ".lintstagedrc.json")
	checkFile(t, tmpDir, "flowkit.json")
}

func TestGenerateNonJS(t *testing.T) {
	cfg := &config.Config{
		ProjectName:   "go-service",
		MainBranch:    "main",
		Language:      "en",
		WorkflowStyle: "github-flow",
		Stack:         "go",
		Features: config.FeaturesConfig{
			CI:          true,
			PRCheck:     false,
			CommitHooks: true,
			WorkflowDoc: true,
		},
		Commands: config.CommandsConfig{
			Install: "go mod download",
			Dev:     "go run .",
			Build:   "go build .",
			Lint:    "go vet ./...",
		},
	}

	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(origDir)

	if err := GenerateAll(cfg, true); err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	checkFile(t, tmpDir, "WORKFLOW.md")
	checkFile(t, tmpDir, ".github/workflows/ci.yml")
	checkFile(t, tmpDir, ".githooks/pre-commit")
	checkFile(t, tmpDir, ".githooks/commit-msg")
	checkFileNotExist(t, tmpDir, "commitlint.config.js")
	checkFileNotExist(t, tmpDir, ".lintstagedrc.json")
	checkFile(t, tmpDir, "flowkit.json")
}

func checkFile(t *testing.T, dir, name string) {
	t.Helper()
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("expected file %s does not exist", path)
	}
}

func checkFileNotExist(t *testing.T, dir, name string) {
	t.Helper()
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); err == nil {
		t.Errorf("unexpected file %s exists", path)
	}
}
