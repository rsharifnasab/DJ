package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func prepareSomeFiles() ([]string, string) {
	tmp := MakeTempfolder()
	//defer os.RemoveAll(tmp)
	files := []string{tmp + "/example.txt", tmp + "/data.txt"}

	for _, name := range files {
		err := os.WriteFile(name, []byte("dummy text to write in "+name), 0666)
		cobra.CheckErr(err)
	}
	return files, tmp
}

// With relative paths (e.g. a relative --result directory) ZipDir
// used to chdir first and then create the zip relative to the NEW
// working directory, which failed with "no such file or directory".
func TestZipDirRelativePaths(t *testing.T) {
	origCwd, err := os.Getwd()
	assert.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, os.Chdir(origCwd)) })

	tmp := t.TempDir()
	assert.NoError(t, os.Chdir(tmp))

	res := "result"
	assert.NoError(t, os.MkdirAll(res, 0755))
	assert.NoError(t, os.WriteFile(filepath.Join(res, "result.txt"), []byte("score 100"), 0644))

	assert.NoError(t, ZipDir(filepath.Join(res, "submission.zip"), res))
	assert.FileExists(t, filepath.Join(tmp, "result", "submission.zip"))
}

func TestZipFiles(t *testing.T) {
	files, dir := prepareSomeFiles()
	output := dir + "/done.zip"
	err := ZipFiles(output, files)
	assert.NoError(t, err)

	assert.FileExists(t, output)
	defer os.RemoveAll(output)
}

func TestZipDir(t *testing.T) {
	_, dir := prepareSomeFiles()
	output := dir + "/done.zip"
	err := ZipDir(output, dir)
	assert.NoError(t, err)
	assert.FileExists(t, output)
	defer os.RemoveAll(output)
}

func TestUnzip(t *testing.T) {
	files, dir := prepareSomeFiles()
	output := dir + "/done.zip"
	err := ZipDir(output, dir)
	assert.NoError(t, err)
	assert.FileExists(t, output)
	defer os.RemoveAll(output)

	unzipped := MakeTempfolder()
	defer os.RemoveAll(unzipped)
	fileList, err := Unzip(output, unzipped)
	assert.NoError(t, err)
	assert.Len(t, fileList, len(files))
	assert.FileExists(t, unzipped+"/example.txt")
	assert.FileExists(t, unzipped+"/data.txt")
}

func TestISZip(t *testing.T) {
	files, dir := prepareSomeFiles()
	output := dir + "/done.zip"
	err := ZipFiles(output, files)
	assert.NoError(t, err)

	assert.FileExists(t, output)
	res, err := IsZip(output)
	assert.NoError(t, err)
	assert.True(t, res)
	defer os.RemoveAll(output)
}
