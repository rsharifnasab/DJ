package util

import (
	"os"
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
