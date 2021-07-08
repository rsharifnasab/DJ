package util

import (
	"fmt"
	"os"
)

func CheckFileExists(path string) error {
	if stat, err := os.Stat(path); err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("%v is a directory", path)
	} else {
		return nil
	}
}
