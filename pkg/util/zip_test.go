package util

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestZip1(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)
	files := []string{tmp + "/example.csv", tmp + "/data.csv"}
	output := tmp + "/done.zip"

	for _, name := range files {
		err := os.WriteFile(name, []byte("dummy text to write"), 0666)
		cobra.CheckErr(err)
	}

	err := ZipFiles(output, files)
	cobra.CheckErr(err)

	assert.FileExists(t, output)
	err = os.RemoveAll(output)
	cobra.CheckErr(err)
}
