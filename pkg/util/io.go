package util

import (
	"errors"
	"io"
	"os"

	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

func MakeTempfolder() string {
	tmpFolder, err := os.MkdirTemp("", "dj-tmp-*")
	cobra.CheckErr(err)
	return tmpFolder
}

func CopyDir(srcPath, destPath string) {
	err := cp.Copy(srcPath, destPath)
	cobra.CheckErr(err)
}

func IsDirEmpty(dir string) bool {
	f, err := os.Open(dir)
	cobra.CheckErr(err)
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}
	cobra.CheckErr(err)
	return false
}

func IsDirExists(dir string) bool {
	_, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		cobra.CheckErr(err)
		return true
	}
}
