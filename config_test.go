package main

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func _setup_directory() {
	os.Mkdir(".test", 0755)
	os.Chdir(".test")
}

func _tear_directory() {
	os.Chdir("..")
	os.RemoveAll(".test")
}

func TestConfig(t *testing.T) {
	_setup_directory()
	createInitialFiles(".")
	t.Run("createInitialFiles", func (t *testing.T)  {
		assert.FileExists(t, ".iplan.yaml")
		assert.FileExists(t, ".gitignore")
		assert.DirExists(t,".templates")
	})
	t.Run("initConfig", func (t *testing.T) {
		initConfig()
		expected, _ := filepath.Abs(".iplan.yaml")
		assert.Equal(t, expected, viper.ConfigFileUsed() )
	})
	_tear_directory()
}
