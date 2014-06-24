package main

import (
	"os"
	"os/exec"
)

func main() {
	appDir := os.Args[1]

	err := os.Chdir(appDir)
	if err != nil {
		os.Exit(1)
	}

	cmd := exec.Command("powershell.exe", append([]string{"-Command"}, os.Args[2:]...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
