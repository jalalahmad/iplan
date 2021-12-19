package iplan

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func InitEmptyGit(basePath string) (repository *git.Repository, err error) {
	return git.PlainInit(basePath, false)
}

func AddAndCommitFile(repository *git.Repository, filePath string) (plumbing.Hash, error) {
	var h plumbing.Hash
	w, err := repository.Worktree()
	if err != nil {
		return h, err
	}
	fmt.Println(w.Status())
	w.Add(filePath)
	fmt.Println(w.Status())
	h, err = w.Commit(fmt.Sprintf(
		"Add / Update %s",
		filePath,
	), &git.CommitOptions{})
	fmt.Println(w.Status())

	return h, err
}
