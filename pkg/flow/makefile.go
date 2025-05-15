package flow

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/osbertngok/famg/pkg/cmd"
)

// CreateMakefileResult represents the outcome of makefile creation
type CreateMakefileResult int

const (
	// MakefileCreated indicates successful makefile creation
	MakefileCreated CreateMakefileResult = iota
	// MakefileExists indicates makefile already exists
	MakefileExists
	// MakefileError indicates other errors
	MakefileError
)

// String returns a human-readable description of the CreateMakefileResult
func (r CreateMakefileResult) String() string {
	switch r {
	case MakefileCreated:
		return "Makefile created successfully"
	case MakefileExists:
		return "Makefile already exists"
	case MakefileError:
		return "Error creating Makefile"
	default:
		return "Unknown result"
	}
}

// CreateMakefile creates a Makefile in the specified directory using the template
func CreateMakefile(config cmd.Config) CreateMakefileResult {
	makefilePath := filepath.Join(config.Path, "Makefile")

	// Check if Makefile already exists
	if _, err := os.Stat(makefilePath); err == nil {
		return MakefileExists
	}

	// Parse the template
	tmpl, err := template.ParseFiles("pkg/flow/templates/Makefile.tmpl")
	if err != nil {
		return MakefileError
	}

	// Create the Makefile
	file, err := os.Create(makefilePath)
	if err != nil {
		return MakefileError
	}
	defer file.Close()

	// Execute the template
	if err := tmpl.Execute(file, config); err != nil {
		os.Remove(makefilePath) // Clean up on error
		return MakefileError
	}

	return MakefileCreated
}
