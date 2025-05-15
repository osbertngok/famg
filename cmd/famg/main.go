package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readReadme() (string, error) {
	// Get the current working directory (project root)
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}

	// Try to read README.md from the project root
	readmePath := filepath.Join(wd, "README.md")
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return "", fmt.Errorf("failed to read README.md: %v", err)
	}

	return string(content), nil
}

func main() {
	// If no arguments are provided or if help is requested
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		helpMsg, err := readReadme()
		if err != nil {
			fmt.Printf("Error reading help message: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(helpMsg)
		os.Exit(0)
	}

	// TODO: Add other command handling here
	fmt.Println("Command not implemented yet")
}
