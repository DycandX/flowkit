package generator

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/DycandX/flowkit/internal/config"
	"github.com/DycandX/flowkit/internal/detector"
)

func TestE2ENextJS(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"next":"14"}}`), 0644)

	stack := string(detector.Detect())
	if stack != "next" {
		t.Fatalf("Detect() = %q, want %q", stack, "next")
	}

	cfg := &config.Config{
		ProjectName:   "test-app",
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
		Commands: config.CommandsConfig{
			Install: "npm install",
			Dev:     "npm run dev",
			Build:   "npm run build",
			Lint:    "npm run lint",
		},
	}

	if err := GenerateAll(cfg, true); err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	mustExist(t, "WORKFLOW.md")
	mustExist(t, ".github/workflows/ci.yml")
	mustExist(t, ".github/workflows/pr-check.yml")
	mustExist(t, ".husky/pre-commit")
	mustExist(t, ".husky/commit-msg")
	mustExist(t, "commitlint.config.js")
	mustExist(t, ".lintstagedrc.json")
	mustExist(t, "flowkit.json")

	mustNotExist(t, ".githooks/pre-commit")

	data, err := os.ReadFile("flowkit.json")
	if err != nil {
		t.Fatal(err)
	}
	var saved config.Config
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatalf("flowkit.json invalid JSON: %v", err)
	}
	if saved.ProjectName != "test-app" {
		t.Errorf("flowkit.json ProjectName = %q, want %q", saved.ProjectName, "test-app")
	}
}

func TestE2EGo(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("go.mod", []byte("module test"), 0644)

	stack := string(detector.Detect())
	if stack != "go" {
		t.Fatalf("Detect() = %q, want %q", stack, "go")
	}

	cfg := &config.Config{
		ProjectName:   "go-service",
		MainBranch:    "main",
		WorkflowStyle: "gitflow",
		Stack:         stack,
		Features: config.FeaturesConfig{
			CI:          true,
			PRCheck:     true,
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

	if err := GenerateAll(cfg, true); err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	mustExist(t, "WORKFLOW.md")
	mustExist(t, ".github/workflows/ci.yml")
	mustExist(t, ".githooks/pre-commit")
	mustExist(t, ".githooks/commit-msg")
	mustExist(t, "flowkit.json")

	mustNotExist(t, ".husky/pre-commit")
	mustNotExist(t, ".husky/commit-msg")
	mustNotExist(t, "commitlint.config.js")
	mustNotExist(t, ".lintstagedrc.json")
}

func TestE2ESkipExisting(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"next":"14"}}`), 0644)

	cfg := &config.Config{
		ProjectName:   "test-app",
		MainBranch:    "master",
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

	if err := GenerateAll(cfg, false); err != nil {
		t.Fatalf("first GenerateAll failed: %v", err)
	}

	mtime1 := modTime(t, "WORKFLOW.md")

	if err := GenerateAll(cfg, false); err != nil {
		t.Fatalf("second GenerateAll failed: %v", err)
	}

	mtime2 := modTime(t, "WORKFLOW.md")
	if mtime1 != mtime2 {
		t.Error("WORKFLOW.md should not have been modified on second run (force=false)")
	}
}

func TestE2EForceOverwrite(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"next":"14"}}`), 0644)

	cfg := &config.Config{
		ProjectName:   "test-app",
		MainBranch:    "master",
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

	if err := GenerateAll(cfg, true); err != nil {
		t.Fatalf("first GenerateAll failed: %v", err)
	}

	os.Remove("WORKFLOW.md")
	mustNotExist(t, "WORKFLOW.md")

	if err := GenerateAll(cfg, true); err != nil {
		t.Fatalf("second GenerateAll failed: %v", err)
	}

	mustExist(t, "WORKFLOW.md")
}

func mustExist(t *testing.T, name string) {
	t.Helper()
	if _, err := os.Stat(name); os.IsNotExist(err) {
		t.Errorf("expected file %s does not exist", name)
	}
}

func mustNotExist(t *testing.T, name string) {
	t.Helper()
	if _, err := os.Stat(name); err == nil {
		t.Errorf("unexpected file %s exists", name)
	}
}

func modTime(t *testing.T, name string) int64 {
	t.Helper()
	fi, err := os.Stat(name)
	if err != nil {
		t.Fatal(err)
	}
	return fi.ModTime().UnixNano()
}
