package cmd

import "github.com/spf13/cobra"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize directory for iplan use",
	Long: `Sets up config file with sample story and make directory versionable.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
