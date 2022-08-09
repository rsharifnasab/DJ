package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func LogToResult(basePath, midPath, fileName string, data string) {
	var logDir string
	if midPath == "" {
		logDir = basePath
	} else {
		logDir = basePath + "/" + midPath
	}

	if _, err := os.Stat(logDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(logDir, os.ModePerm)
		cobra.CheckErr(err)
	}

	logFile := fmt.Sprintf("%s/%s.log", logDir, fileName)
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cobra.CheckErr(err)
	_, err = f.Write([]byte(data))
	cobra.CheckErr(err)
	err = f.Close()
	cobra.CheckErr(err)
}
