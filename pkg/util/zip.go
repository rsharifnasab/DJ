package util

import (
	"archive/zip"
	"io"
	"os"
)

func ZipDirNaive(filename string, dirs []string) error {
	list := make([]string, 0)
	for _, dir := range dirs {
		dirList, err := WalkDir(dir)
		if err != nil {
			return err
		}
		list = append(list, dirList...)
	}
	return ZipFiles(filename, list)
}

func ZipDir(filename string, dir string) error {
	var err error

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer func(cwd string) {
		err = os.Chdir(dir)
		if err != nil {
			panic(err)
		}
	}(cwd)

	err = os.Chdir(dir)
	if err != nil {
		return err
	}

	list, err := WalkDir(".")
	if err != nil {
		return err
	}

	return ZipFiles(filename, list)
}

func ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
