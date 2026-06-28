package detector

import (
	"encoding/json"
	"os"
	"strings"
)

type Stack string

const (
	StackNext    Stack = "next"
	StackReact   Stack = "react"
	StackVue     Stack = "vue"
	StackNuxt    Stack = "nuxt"
	StackNode    Stack = "node"
	StackRust    Stack = "rust"
	StackGo      Stack = "go"
	StackPython  Stack = "python"
	StackLaravel Stack = "laravel"
	StackUnknown Stack = "unknown"
)

type PkgJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func Detect() Stack {
	if hasFile("package.json") {
		return detectNode()
	}
	if hasFile("Cargo.toml") {
		return StackRust
	}
	if hasFile("go.mod") {
		return StackGo
	}
	if hasFile("pyproject.toml") || hasFile("requirements.txt") {
		return StackPython
	}
	if hasFile("composer.json") {
		return StackLaravel
	}
	return StackUnknown
}

func detectNode() Stack {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return StackNode
	}
	var pkg PkgJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return StackNode
	}
	all := make(map[string]string)
	for k, v := range pkg.Dependencies {
		all[k] = v
	}
	for k, v := range pkg.DevDependencies {
		all[k] = v
	}
	if hasDep(all, "next") {
		return StackNext
	}
	if hasDep(all, "nuxt") {
		return StackNuxt
	}
	if hasDep(all, "vue") {
		return StackVue
	}
	if hasDep(all, "react") {
		return StackReact
	}
	return StackNode
}

func hasDep(deps map[string]string, name string) bool {
	for k := range deps {
		if strings.EqualFold(k, name) {
			return true
		}
	}
	return false
}

func hasFile(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
