package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
)

const (
	CfgFileName = ".iplan"
	GitIgnoreFileName = ".gitignore"
	SampleCfgFileContent = `---
author: YOUR NAME <you@email.me>
`
	GitIgnoreFileContent = ` #Ignore Setting
.iplan.*
`
)
var cfgFile = ""


func _createFileWithContent(filepath string, content string) error {
	contentBytes := []byte(content)
	return ioutil.WriteFile(filepath, contentBytes, 0644)
}

func createInitialFiles(basePath string) error {
	err := _createFileWithContent(path.Join(basePath, CfgFileName + ".yaml"), SampleCfgFileContent)
	if err != nil {
		return err
	}

	return _createFileWithContent(path.Join(basePath, GitIgnoreFileName), GitIgnoreFileContent)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, _ := os.Getwd()
		viper.AddConfigPath(cwd)
		viper.SetConfigName(CfgFileName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}
