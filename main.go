package main

import (
	"github.com/DycandX/flowkit/cmd"
)

var version = "dev"

func main() {
	cmd.Version = version
	cmd.Execute()
}
