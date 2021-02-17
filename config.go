package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)


var cfgFile = ""


func createInitialFiles(basePath string) error {
	err := RestoreAssets(".","")
	if err!=nil {
		return err
	}

	return os.Rename("gitignore", ".gitignore")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, _ := os.Getwd()
		viper.AddConfigPath(cwd)
		viper.SetConfigName(".iplan")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}
