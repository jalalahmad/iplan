package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var CfgFile = ""

func CreateInitialFiles(basePath string) error {
	err := RestoreAssets(".", "")
	if err != nil {
		return err
	}

	return os.Rename("gitignore", ".gitignore")
}

func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
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
