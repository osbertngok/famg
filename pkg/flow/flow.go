package flow

import (
	"fmt"

	_ "embed"

	"github.com/osbertngok/famg/pkg/cmd"
)

// MainFlow orchestrates the folder creation, git repository initialization, .gitignore population, and Makefile creation
func MainFlow(config cmd.Config) {
	if createFolderResult := CreateFolder(config); createFolderResult != FolderCreated {
		fmt.Println(createFolderResult.String())
		return
	}

	if createGitRepoResult := CreateGitRepo(config); createGitRepoResult != GitRepoCreated {
		fmt.Println(createGitRepoResult.String())
		return
	}

	if populateGitignoreResult := PopulateGitignore(config); populateGitignoreResult != GitignorePopulated {
		fmt.Println(populateGitignoreResult.String())
		return
	}

	if createMakefileResult := CreateMakefile(config); createMakefileResult != MakefileCreated {
		fmt.Println(createMakefileResult.String())
		return
	}
}
