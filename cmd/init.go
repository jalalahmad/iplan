package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize directory for iplan use",
	Long:  `Sets up config file with sample story and make directory versionable.`,
	Run: func(cmd *cobra.Command, args []string) {
		initGit()
		createConfigFile()
	},
}

func initGit() (ok bool) {
	path, err := os.Getwd()
	if err != nil {
		return false
	}

	_, err = git.PlainInit(path, false)
	if err != nil {
		return false
	}

	return true
}

func createConfigFile() (ok bool) {
	d1 := []byte(`---
author: YOUR NAME <you@email.me>
`)
	err := ioutil.WriteFile("./.iplan.yaml", d1, 0644)
	if err != nil {
		return false
	}

	return true
}
