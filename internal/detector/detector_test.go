package detector

import (
	"os"
	"os/exec"
	"testing"
)

func TestDetectNextJS(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"next":"14"}}`), 0644)
	if got := Detect(); got != StackNext {
		t.Errorf("Detect() = %v, want %v", got, StackNext)
	}
}

func TestDetectReact(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"react":"18"}}`), 0644)
	if got := Detect(); got != StackReact {
		t.Errorf("Detect() = %v, want %v", got, StackReact)
	}
}

func TestDetectVue(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"vue":"3"}}`), 0644)
	if got := Detect(); got != StackVue {
		t.Errorf("Detect() = %v, want %v", got, StackVue)
	}
}

func TestDetectNuxt(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{"dependencies":{"nuxt":"3"}}`), 0644)
	if got := Detect(); got != StackNuxt {
		t.Errorf("Detect() = %v, want %v", got, StackNuxt)
	}
}

func TestDetectNode(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("package.json", []byte(`{}`), 0644)
	if got := Detect(); got != StackNode {
		t.Errorf("Detect() = %v, want %v", got, StackNode)
	}
}

func TestDetectGo(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("go.mod", []byte("module test"), 0644)
	if got := Detect(); got != StackGo {
		t.Errorf("Detect() = %v, want %v", got, StackGo)
	}
}

func TestDetectRust(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("Cargo.toml", []byte("[package]\nname = \"test\""), 0644)
	if got := Detect(); got != StackRust {
		t.Errorf("Detect() = %v, want %v", got, StackRust)
	}
}

func TestDetectPython(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("requirements.txt", []byte(""), 0644)
	if got := Detect(); got != StackPython {
		t.Errorf("Detect() = %v, want %v", got, StackPython)
	}
}

func TestDetectLaravel(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	os.WriteFile("composer.json", []byte("{}"), 0644)
	if got := Detect(); got != StackLaravel {
		t.Errorf("Detect() = %v, want %v", got, StackLaravel)
	}
}

func TestDetectUnknown(t *testing.T) {
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	exec.Command("git", "init").Run()
	if got := Detect(); got != StackUnknown {
		t.Errorf("Detect() = %v, want %v", got, StackUnknown)
	}
}
