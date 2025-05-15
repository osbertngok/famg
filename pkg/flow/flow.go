package flow

import (
	"fmt"
	"os"
	"os/exec"
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

// CreateGitRepoResult represents the outcome of git repository initialization
type CreateGitRepoResult int

const (
	// GitRepoCreated indicates successful git repository creation
	GitRepoCreated CreateGitRepoResult = iota
	// GitRepoExists indicates git repository already exists
	GitRepoExists
	// GitNotInstalled indicates git is not installed
	GitNotInstalled
	// GitInitError indicates other git initialization errors
	GitInitError
)

// String returns a human-readable description of the CreateGitRepoResult
func (r CreateGitRepoResult) String() string {
	switch r {
	case GitRepoCreated:
		return "Git repository created successfully"
	case GitRepoExists:
		return "Git repository already exists"
	case GitNotInstalled:
		return "Git is not installed"
	case GitInitError:
		return "Error initializing git repository"
	default:
		return "Unknown result"
	}
}

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

// CreateGitRepo initializes a git repository in the specified folder
// Returns a CreateGitRepoResult indicating the outcome
func CreateGitRepo(config cmd.Config) CreateGitRepoResult {
	// Ensure the path is absolute
	absPath, err := filepath.Abs(config.Path)
	if err != nil {
		return GitInitError
	}

	// Check if git is installed
	if _, err := exec.LookPath("git"); err != nil {
		return GitNotInstalled
	}

	// Check if .git directory already exists
	if _, err := os.Stat(filepath.Join(absPath, ".git")); err == nil {
		return GitRepoExists
	}

	// Initialize git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = absPath
	if err := cmd.Run(); err != nil {
		return GitInitError
	}

	return GitRepoCreated
}

func MainFlow(config cmd.Config) {
	if createFolderResult := CreateFolder(config); createFolderResult != FolderCreated {
		fmt.Println(createFolderResult.String())
		return
	}

	if createGitRepoResult := CreateGitRepo(config); createGitRepoResult != GitRepoCreated {
		fmt.Println(createGitRepoResult.String())
		return
	}
}
