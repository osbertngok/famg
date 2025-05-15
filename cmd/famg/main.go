package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/osbertngok/famg/pkg/cmd"
	"github.com/osbertngok/famg/pkg/flow"
	"github.com/osbertngok/famg/pkg/help"
	"github.com/spf13/cobra"
)

var (
	configFile string
	parentPath string
	name       string
	fullName   string
)

var rootCmd = &cobra.Command{
	Use:   "famg",
	Short: "FAMG - File and Git Management Tool",
	Long:  help.HelpText,
	Run: func(cobraCmd *cobra.Command, args []string) {
		// If config file is provided, load config from file
		if configFile != "" {
			// TODO: Implement config file loading
			fmt.Println("Config file loading not implemented yet")
			return
		}

		// If command line arguments are provided, use them
		if parentPath != "" && name != "" && fullName != "" {
			absPath, err := filepath.Abs(parentPath)
			if err != nil {
				fmt.Printf("Error: invalid parent path: %v\n", err)
				os.Exit(1)
			}

			config := cmd.Config{
				Path:       filepath.Join(absPath, name),
				Name:       name,
				FullName:   fullName,
				ParentPath: absPath,
			}

			flow.MainFlow(config)
			return
		}

		// If no valid arguments are provided, show help
		cobraCmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config-file", "", "Path to the config file")
	rootCmd.Flags().StringVar(&parentPath, "parent-path", "", "Parent path for the new folder")
	rootCmd.Flags().StringVar(&name, "name", "", "Name of the folder to be created")
	rootCmd.Flags().StringVar(&fullName, "fullname", "", "Full name of the folder to be created")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
