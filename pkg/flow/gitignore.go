package flow

import (
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/osbertngok/famg/pkg/cmd"
)

//go:embed gitignore-sample
var gitignoreContent string

// PopulateGitignoreResult represents the outcome of populating the .gitignore file
type PopulateGitignoreResult int

const (
	// GitignorePopulated indicates successful population of .gitignore
	GitignorePopulated PopulateGitignoreResult = iota
	// GitignoreExists indicates .gitignore already exists
	GitignoreExists
	// GitignoreError indicates other errors
	GitignoreError
)

// String returns a human-readable description of the PopulateGitignoreResult
func (r PopulateGitignoreResult) String() string {
	switch r {
	case GitignorePopulated:
		return ".gitignore populated successfully"
	case GitignoreExists:
		return ".gitignore already exists"
	case GitignoreError:
		return "Error populating .gitignore"
	default:
		return "Unknown result"
	}
}

func PopulateGitignore(config cmd.Config) PopulateGitignoreResult {
	gitignore := filepath.Join(config.Path, ".gitignore")
	if _, err := os.Stat(gitignore); os.IsNotExist(err) {
		// Write the embedded .gitignore content to the file
		if err := os.WriteFile(gitignore, []byte(gitignoreContent), 0644); err != nil {
			fmt.Printf("Error writing .gitignore: %v\n", err)
			return GitignoreError
		}
		return GitignorePopulated
	}
	return GitignoreExists
}
