package main

import (
	"os/exec"
)

// Build ...
func Build() {
	cmd := exec.Command("go", "build", path)
	cmd.Dir = path[:strings.LastIndex(path, `/`)]
}
