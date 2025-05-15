package flow

import (
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/osbertngok/famg/pkg/cmd"
	"go.uber.org/zap"
)

// CreatePyvenvResult represents the outcome of pyvenv.cfg creation
type CreatePyvenvResult int

const (
	PyvenvCreated CreatePyvenvResult = iota
	PyvenvExists
	PyvenvError
)

func (r CreatePyvenvResult) String() string {
	switch r {
	case PyvenvCreated:
		return "pyvenv.cfg created successfully"
	case PyvenvExists:
		return "pyvenv.cfg already exists"
	case PyvenvError:
		return "Error creating pyvenv.cfg"
	default:
		return "Unknown result"
	}
}

func CreatePyvenv(config cmd.Config) CreatePyvenvResult {
	venvDir := filepath.Join(config.Path, ".ve3")
	if err := os.MkdirAll(venvDir, 0755); err != nil {
		logger.Error("Failed to create .ve3 directory", zap.Error(err))
		return PyvenvError
	}
	pyvenvPath := filepath.Join(venvDir, "pyvenv.cfg")
	if _, err := os.Stat(pyvenvPath); err == nil {
		logger.Info("pyvenv.cfg already exists", zap.String("path", pyvenvPath))
		return PyvenvExists
	}
	tmpl, err := template.ParseFiles("pkg/flow/templates/pyvenv.cfg.tmpl")
	if err != nil {
		logger.Error("Failed to parse pyvenv.cfg template", zap.Error(err))
		return PyvenvError
	}
	file, err := os.Create(pyvenvPath)
	if err != nil {
		logger.Error("Failed to create pyvenv.cfg", zap.String("path", pyvenvPath), zap.Error(err))
		return PyvenvError
	}
	defer file.Close()
	if err := tmpl.ExecuteTemplate(file, "pyvenv.cfg.tmpl", config); err != nil {
		logger.Error("Failed to execute pyvenv.cfg template", zap.Error(err))
		os.Remove(pyvenvPath)
		return PyvenvError
	}
	// Force add and commit the pyvenv.cfg file
	cmd := exec.Command("git", "add", "-f", ".ve3/pyvenv.cfg")
	cmd.Dir = config.Path
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to add pyvenv.cfg", zap.Error(err))
		return PyvenvError
	}
	cmd = exec.Command("git", "commit", "-m", "feat(init): add pyvenv.cfg")
	cmd.Dir = config.Path
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to commit pyvenv.cfg", zap.Error(err))
		return PyvenvError
	}
	logger.Info("pyvenv.cfg created successfully", zap.String("path", pyvenvPath))
	return PyvenvCreated
}
