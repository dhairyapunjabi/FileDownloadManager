package model

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type Download interface {
	DownloadUrls()
}

type SerialDownload struct {
	Urls []string
}

type ConcurrentDownload struct {
	Urls []string
}

func (serial_download SerialDownload) DownloadUrls() {
	for _, url := range serial_download.Urls {
		_ = DownloadFile(url)
	}
}

func (concurrent_download ConcurrentDownload) DownloadUrls() {
}

type DownloadId struct {
	Id string
}

func DownloadFile(url string) error {
	filepath := "/tmp" + "/" + GenerateUuid()
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GenerateUuid() string {
	id := uuid.New()
	return id.String()
}
