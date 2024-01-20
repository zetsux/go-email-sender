package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const (
	DEFAULT_FILE_NAME      = "file"
	DEFAULT_FILE_DIRECTORY = "files"
	DEFAULT_DATA_NAME      = "data.csv"
	DEFAULT_DATA_DIRECTORY = "data"
)

func DownloadFileFromURL(url string, fileName string) (err error) {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	fmt.Println(r.Status)

	if r.StatusCode != http.StatusOK {
		return err
	}

	var fileExt string
	contentType := r.Header.Get("Content-Type")
	exts, err := mime.ExtensionsByType(contentType)
	if err != nil || len(exts) == 0 {
		fileExt = filepath.Ext(url)
	} else {
		fileExt = exts[0]
	}

	err = os.MkdirAll(DEFAULT_FILE_DIRECTORY, os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(DEFAULT_FILE_DIRECTORY + "/" + fileName + fileExt)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		panic(err)
	}
	return nil
}
