package util

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestDirEmpty(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	assert.True(t, IsDirEmpty(tmp))
}

func TestDirNotEmpty(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	err := os.WriteFile(tmp+"/dummy.txt", []byte("dummpy text to write"), 0666)
	cobra.CheckErr(err)

	assert.False(t, IsDirEmpty(tmp))
}

func TestDirExists(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	assert.True(t, IsDirExists(tmp))
}

func TestDirNotExists(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)

	assert.False(t, IsDirExists(tmp+"/subfolder"))
}
