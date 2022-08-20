package util

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewLoggerEmptyPath(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	l := NewLogger(tmp)
	assert.Equal(t, false, l.nonEmptyPath)
}

func TestNewLoggerNonEmptyPath(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)

	err := os.WriteFile(tmp+"/dummy.txt", []byte("dummy text to write"), 0666)
	cobra.CheckErr(err)

	l := NewLogger(tmp)
	assert.Equal(t, true, l.nonEmptyPath)
	assert.NotEqual(t, tmp, l.BasePath)
	assert.Equal(t, tmp, l.userBasePath)
}

func TestNewLoggerNonExistingPath(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	l := NewLogger(tmp + "/non_existing_path")
	assert.Equal(t, true, l.nonExistingDir)
	assert.DirExists(t, tmp+"/non_existing_path")
}

func TestLogToRoot(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	l := NewLogger(tmp)
	l.LogTo("", "file", "dummy text")
	assert.FileExists(t, tmp+"/file.log")
}

func TestLogToFolder(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	l := NewLogger(tmp)
	l.LogTo("dir", "file", "dummy text")
	assert.FileExists(t, tmp+"/dir/file.log")
}
