package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
)

func zipFiles(output, input string) error {
	dest, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

	return addToZips(input, zipWriter)
}

func addToZips(path string, zipWriter *zip.Writer) error {
	inputInfo, err := os.Stat(path)

	if err != nil {
		return err
	}

	if inputInfo.IsDir() {
		fileInfos, err := ioutil.ReadDir(inputInfo.Name())
		if err != nil {
			return err
		}

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				err := addToZips(fileInfo.Name(), zipWriter)

				if err != nil {
					return err
				}
			} else {
				addToZip(fileInfo.Name(), zipWriter)

				if err != nil {
					return err
				}
			}
		}
	} else {
		err := addToZip(inputInfo.Name(), zipWriter)

		if err != nil {
			return err
		}
	}

	return nil
}

func addToZip(filename string, zipWriter *zip.Writer) error {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	writer, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, src)
	if err != nil {
		return err
	}

	return nil
}
