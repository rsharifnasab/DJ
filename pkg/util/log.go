package util

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Logger struct {
	userBasePath   string
	BasePath       string
	nonEmptyPath   bool
	nonExistingDir bool
}

func NewLogger(basePath string) *Logger {
	logger := &Logger{
		BasePath:       basePath,
		userBasePath:   basePath,
		nonEmptyPath:   false,
		nonExistingDir: false,
	}
	if !IsDirExists(basePath) {
		err := os.Mkdir(basePath, os.ModePerm)
		cobra.CheckErr(err)
		logger.nonExistingDir = true
	}
	if !IsDirEmpty(basePath) {
		createdPath, err := os.MkdirTemp(basePath, "dj-result-*")
		cobra.CheckErr(err)
		logger.BasePath = createdPath
		logger.nonEmptyPath = true
	}
	return logger
}

func (logger *Logger) LogTo(midPath, fileName string, data string) {
	var logDir string
	if midPath == "" {
		logDir = logger.BasePath
	} else {
		logDir = logger.BasePath + "/" + midPath
	}

	if !IsDirExists(logDir) {
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
