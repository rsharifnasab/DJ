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
	err := os.WriteFile(tmp+"/dummy.txt", []byte("dummy text to write"), 0666)
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

func TestWalk(t *testing.T) {
	var err error

	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	err = os.MkdirAll(tmp+"/a/b/c/d", 0777)
	assert.NoError(t, err)

	err = os.WriteFile(tmp+"/r1.txt", []byte("dummy text to write"), 0666)
	assert.NoError(t, err)

	err = os.WriteFile(tmp+"/a/a1.txt", []byte("dummy text to write"), 0666)
	assert.NoError(t, err)
	err = os.WriteFile(tmp+"/a/a2.txt", []byte("dummy text to write"), 0666)
	assert.NoError(t, err)
	err = os.WriteFile(tmp+"/a/b/c/d/d1.txt", []byte("dummy text to write"), 0666)
	assert.NoError(t, err)

	list, err := WalkDir(tmp)
	assert.Len(t, list, 4)
}
