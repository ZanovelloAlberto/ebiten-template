//go:build example
// +build example

package main

import "os/exec"

func main() {
	buildLinux()
}

func buildLinux() {
	exec.Command("go", "build", "-o", "build/linux", "./src/main.go").Start()
}
