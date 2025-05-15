package flow

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}
	tmpl, err := template.New("Makefile").Funcs(funcMap).ParseFiles("pkg/flow/templates/Makefile.tmpl")
	if err != nil {
		logger.Error("Failed to parse Makefile template",
			zap.Error(err))
		return MakefileError
	}

	// Log the config values
	logger.Info("Creating Makefile with config",
		zap.String("Name", config.Name),
		zap.String("Path", config.Path))

	// Log the template name
	logger.Info("Template name",
		zap.String("name", tmpl.Name()))

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
	if err := tmpl.ExecuteTemplate(file, "Makefile.tmpl", config); err != nil {
		logger.Error("Failed to execute template",
			zap.Error(err))
		os.Remove(makefilePath) // Clean up on error
		return MakefileError
	}

	// Force add and commit the Makefile
	cmd := exec.Command("git", "add", "-f", "Makefile")
	cmd.Dir = config.Path
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to add Makefile",
			zap.Error(err))
		return MakefileError
	}
	cmd = exec.Command("git", "commit", "-m", "feat(init): add Makefile")
	cmd.Dir = config.Path
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to commit Makefile",
			zap.Error(err))
		return MakefileError
	}

	logger.Info("Makefile created successfully",
		zap.String("path", makefilePath))
	return MakefileCreated
}
