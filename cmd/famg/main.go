package main

import (
	"fmt"
	"os"

	"github.com/osbertngok/famg/pkg/cmd"
	"github.com/osbertngok/famg/pkg/flow"
	"github.com/osbertngok/famg/pkg/help"
)

func main() {
	// If no arguments are provided or if help is requested
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(help.HelpText)
		os.Exit(0)
	}

	// Create config from command line arguments
	config := cmd.Config{
		Path: os.Args[1],
	}

	// Execute the main flow
	flow.MainFlow(config)
}
