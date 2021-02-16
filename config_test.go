package main

import (
	"os"
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
	assert.FileExists(t, ".iplan.yaml")
	assert.FileExists(t, ".gitignore")
	initConfig()
	_tear_directory()

}
