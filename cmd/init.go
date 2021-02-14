package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize directory for iplan use",
	Long:  `Sets up config file with sample story and make directory versionable.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		createConfigFile()
	},
}

func createConfigFile() {
	d1 := []byte(`---
author: YOUR NAME <you@email.me>
`)
	err := ioutil.WriteFile("./.iplan.yaml", d1, 0644)
	cobra.CheckErr(err)
}
