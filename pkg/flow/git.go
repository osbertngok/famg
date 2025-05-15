package flow

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/osbertngok/famg/pkg/cmd"
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
