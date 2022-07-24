package util

import (
	"io/ioutil"

	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

func MakeTempfolder() string {
	tmpFolder, err := ioutil.TempDir("", "dj-tmp-*")
	cobra.CheckErr(err)
	return tmpFolder
}
func CopyDir(srcPath, destPath string) {
	err := cp.Copy(srcPath, destPath)
	cobra.CheckErr(err)
}
