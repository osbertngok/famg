package flow

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/osbertngok/famg/pkg/cmd"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
}

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
		logger.Info("Makefile already exists",
			zap.String("path", makefilePath))
		return MakefileExists
	}

	// Parse the template
	tmpl, err := template.ParseFiles("pkg/flow/templates/Makefile.tmpl")
	if err != nil {
		logger.Error("Failed to parse Makefile template",
			zap.Error(err))
		return MakefileError
	}

	// Create the Makefile
	file, err := os.Create(makefilePath)
	if err != nil {
		logger.Error("Failed to create Makefile",
			zap.String("path", makefilePath),
			zap.Error(err))
		return MakefileError
	}
	defer file.Close()

	// Execute the template
	if err := tmpl.Execute(file, config); err != nil {
		logger.Error("Failed to execute template",
			zap.Error(err))
		os.Remove(makefilePath) // Clean up on error
		return MakefileError
	}

	logger.Info("Makefile created successfully",
		zap.String("path", makefilePath))
	return MakefileCreated
}
