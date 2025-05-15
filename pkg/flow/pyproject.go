package flow

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/osbertngok/famg/pkg/cmd"
	"go.uber.org/zap"
)

// CreatePyprojectResult represents the outcome of pyproject.toml creation
type CreatePyprojectResult int

const (
	PyprojectCreated CreatePyprojectResult = iota
	PyprojectExists
	PyprojectError
)

func (r CreatePyprojectResult) String() string {
	switch r {
	case PyprojectCreated:
		return "pyproject.toml created successfully"
	case PyprojectExists:
		return "pyproject.toml already exists"
	case PyprojectError:
		return "Error creating pyproject.toml"
	default:
		return "Unknown result"
	}
}

func CreatePyproject(config cmd.Config) CreatePyprojectResult {
	pyprojectPath := filepath.Join(config.Path, "pyproject.toml")
	if _, err := os.Stat(pyprojectPath); err == nil {
		logger.Info("pyproject.toml already exists", zap.String("path", pyprojectPath))
		return PyprojectExists
	}
	tmpl, err := template.ParseFiles("pkg/flow/templates/pyproject.toml.tmpl")
	if err != nil {
		logger.Error("Failed to parse pyproject.toml template", zap.Error(err))
		return PyprojectError
	}
	file, err := os.Create(pyprojectPath)
	if err != nil {
		logger.Error("Failed to create pyproject.toml", zap.String("path", pyprojectPath), zap.Error(err))
		return PyprojectError
	}
	defer file.Close()
	if err := tmpl.ExecuteTemplate(file, "pyproject.toml.tmpl", config); err != nil {
		logger.Error("Failed to execute pyproject.toml template", zap.Error(err))
		os.Remove(pyprojectPath)
		return PyprojectError
	}
	logger.Info("pyproject.toml created successfully", zap.String("path", pyprojectPath))
	return PyprojectCreated
}
