package command

import (
	"fmt"
	"github.com/jalalahmad/iplan/internal/app/iplan"
	"github.com/jalalahmad/iplan/internal/configs"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iplan",
	Short: "iplan is utility to manage user stories inn plain text",
	Long: `A Fast and Flexible text based user story management utility
built in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize directory for iplan use",
	Long:  `Sets up config file with sample story and make directory versioned.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		r, err := iplan.InitEmptyGit(cwd)
		if err != nil {
			return err
		}

		err = configs.CreateInitialFiles(cwd)
		if err != nil {
			return err
		}

		_, err = iplan.AddAndCommitFile(r, "gitignore")
		return err
	},
}

func init() {
	cobra.OnInitialize(configs.InitConfig)
	rootCmd.PersistentFlags().StringVar(&configs.CfgFile, "config", "", "config file (default is ./.iplan.yaml)")
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
