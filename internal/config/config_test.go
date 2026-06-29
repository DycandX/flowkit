package config

import (
	"encoding/json"
	"testing"
)

func TestConfigMarshal(t *testing.T) {
	cfg := &Config{
		ProjectName:   "test-project",
		MainBranch:    "main",
		Language:      "en",
		WorkflowStyle: "gitflow",
		Stack:         "next",
		Features: FeaturesConfig{
			CI:          true,
			PRCheck:     true,
			CommitHooks: true,
			WorkflowDoc: true,
		},
		Commands: CommandsConfig{
			Install: "npm install",
			Dev:     "npm run dev",
			Build:   "npm run build",
			Lint:    "npm run lint",
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var cfg2 Config
	if err := json.Unmarshal(data, &cfg2); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if cfg2.ProjectName != cfg.ProjectName {
		t.Errorf("ProjectName = %q, want %q", cfg2.ProjectName, cfg.ProjectName)
	}
	if cfg2.MainBranch != cfg.MainBranch {
		t.Errorf("MainBranch = %q, want %q", cfg2.MainBranch, cfg.MainBranch)
	}
	if cfg2.WorkflowStyle != cfg.WorkflowStyle {
		t.Errorf("WorkflowStyle = %q, want %q", cfg2.WorkflowStyle, cfg.WorkflowStyle)
	}
	if cfg2.Stack != cfg.Stack {
		t.Errorf("Stack = %q, want %q", cfg2.Stack, cfg.Stack)
	}
	if cfg2.Features.CI != cfg.Features.CI {
		t.Errorf("Features.CI = %v, want %v", cfg2.Features.CI, cfg.Features.CI)
	}
	if cfg2.Commands.Install != cfg.Commands.Install {
		t.Errorf("Commands.Install = %q, want %q", cfg2.Commands.Install, cfg.Commands.Install)
	}
}

func TestConfigDefaults(t *testing.T) {
	cfg := &Config{}
	if cfg.MainBranch != "" {
		t.Errorf("default MainBranch should be empty, got %q", cfg.MainBranch)
	}
}

func TestConfigFeatures(t *testing.T) {
	f := FeaturesConfig{
		CI:          true,
		PRCheck:     false,
		CommitHooks: true,
		WorkflowDoc: true,
	}
	if !f.CI {
		t.Error("CI should be true")
	}
	if f.PRCheck {
		t.Error("PRCheck should be false")
	}
	if !f.CommitHooks {
		t.Error("CommitHooks should be true")
	}
}

func TestConfigCommands(t *testing.T) {
	c := CommandsConfig{
		Install: "go mod download",
		Dev:     "go run .",
		Build:   "go build .",
		Lint:    "go vet ./...",
	}
	if c.Install != "go mod download" {
		t.Errorf("Install = %q, want %q", c.Install, "go mod download")
	}
}
