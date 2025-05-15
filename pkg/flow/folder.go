package flow

import (
	"os"
	"path/filepath"

	"github.com/osbertngok/famg/pkg/cmd"
)

// CreateFolderResult represents the outcome of folder creation
type CreateFolderResult int

const (
	// FolderCreated indicates successful folder creation
	FolderCreated CreateFolderResult = iota
	// FolderExists indicates the folder already exists
	FolderExists
	// NoPermission indicates insufficient permissions
	NoPermission
	// FolderError indicates other errors
	UnknownFolderError
)

// String returns a human-readable description of the CreateFolderResult
func (r CreateFolderResult) String() string {
	switch r {
	case FolderCreated:
		return "Folder created successfully"
	case FolderExists:
		return "Folder already exists"
	case NoPermission:
		return "Insufficient permissions"
	case UnknownFolderError:
		return "Unknown error creating folder"
	default:
		return "Unknown result"
	}
}

// CreateFolder creates a folder according to the path specified in Config
// Returns a CreateFolderResult indicating the outcome
func CreateFolder(config cmd.Config) CreateFolderResult {
	// Ensure the path is absolute
	absPath, err := filepath.Abs(config.Path)
	if err != nil {
		return UnknownFolderError
	}

	// Check if directory exists
	if _, err := os.Stat(absPath); err == nil {
		return FolderExists
	}

	// Try to create the directory
	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		// Check for permission error
		if os.IsPermission(err) {
			return NoPermission
		}
		return UnknownFolderError
	}

	return FolderCreated
}
