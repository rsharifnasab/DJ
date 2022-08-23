package util

import (
	"errors"
	"io"
	"os"
	"path/filepath"

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

func WalkDir(dir string) ([]string, error) {
	list := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}
func ListDir(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(files))
	for _, fileInfo := range files {
		res = append(res, fileInfo.Name())
	}

	return res, nil
}

func AutoCd(dir string) string {
	if !IsDirExists(dir) {
		return dir
	} else if listDir, err := ListDir(dir); err != nil {
		return dir
	} else if len(listDir) == 1 {
		if IsDirExists(listDir[0]) {
			return AutoCd(dir + "/" + listDir[0])
		} else {
			return dir // handle single file upload
		}
	} else {
		return dir
	}
}
