package prompts

import (
	"testing"
)

func TestDefaultCommandsJS(t *testing.T) {
	cmds := DefaultCommands("next")
	if cmds.Install != "npm install" {
		t.Errorf("Install = %q, want %q", cmds.Install, "npm install")
	}
	if cmds.Lint != "npm run lint" {
		t.Errorf("Lint = %q, want %q", cmds.Lint, "npm run lint")
	}
}

func TestDefaultCommandsGo(t *testing.T) {
	cmds := DefaultCommands("go")
	if cmds.Install != "go mod download" {
		t.Errorf("Install = %q, want %q", cmds.Install, "go mod download")
	}
	if cmds.Build != "go build ." {
		t.Errorf("Build = %q, want %q", cmds.Build, "go build .")
	}
}

func TestDefaultCommandsRust(t *testing.T) {
	cmds := DefaultCommands("rust")
	if cmds.Install != "cargo build" {
		t.Errorf("Install = %q, want %q", cmds.Install, "cargo build")
	}
	if cmds.Lint != "cargo clippy" {
		t.Errorf("Lint = %q, want %q", cmds.Lint, "cargo clippy")
	}
}

func TestDefaultCommandsPython(t *testing.T) {
	cmds := DefaultCommands("python")
	if cmds.Install != "pip install -r requirements.txt" {
		t.Errorf("Install = %q, want %q", cmds.Install, "pip install -r requirements.txt")
	}
	if cmds.Lint != "ruff check ." {
		t.Errorf("Lint = %q, want %q", cmds.Lint, "ruff check .")
	}
}

func TestDefaultCommandsLaravel(t *testing.T) {
	cmds := DefaultCommands("laravel")
	if cmds.Dev != "php artisan serve" {
		t.Errorf("Dev = %q, want %q", cmds.Dev, "php artisan serve")
	}
}

func TestDefaultCommandsUnknown(t *testing.T) {
	cmds := DefaultCommands("unknown")
	if cmds.Install != "npm install" {
		t.Errorf("Install = %q, want %q", cmds.Install, "npm install")
	}
}
