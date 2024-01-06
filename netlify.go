package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runNetlifyDeployCommand() error {
	cmd := exec.Command("netlify", "deploy", "--prod")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Provide newline character as input to accept the default publish directory
	cmd.Stdin = strings.NewReader("\n")

	fmt.Println("Running Netlify deploy command...")

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Netlify deploy completed successfully.")
	return nil
}
