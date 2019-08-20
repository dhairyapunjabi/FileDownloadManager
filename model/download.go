package model

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

type Download interface {
	DownloadUrls() DownloadId
}

type SerialDownload struct {
	Urls []string
}

type ConcurrentDownload struct {
	Urls []string
}

func (serial_download SerialDownload) DownloadUrls() DownloadId {
	downloadId := DownloadId{Id: GenerateUuid()}
	for _, url := range serial_download.Urls {
		_ = DownloadFile(url)
	}
	return downloadId
}

func (concurrent_download ConcurrentDownload) DownloadUrls() DownloadId {
	start_time := time.Now()
	const noOfWorkers = 5
	reqChan := make(chan string)
	var counter uint64
	var noOfUrls uint64
	downloadId := DownloadId{Id: GenerateUuid()}
	noOfUrls = uint64(len(concurrent_download.Urls))
	for i := range currentlyDownloadedFiles {
		delete(currentlyDownloadedFiles, i)
	}
	DownloadStatusMap[downloadId.Id] = StatusResponse{downloadId.Id, start_time.String(), time.Now().String(), "QUEUED", "CONCURRENT", currentlyDownloadedFiles}
	for i := 0; i < noOfWorkers; i++ {
		go Worker(reqChan, noOfUrls, downloadId.Id, &counter)
	}

	go func() {
		for _, url := range concurrent_download.Urls {
			reqChan <- url
		}
		DownloadStatusMap[downloadId.Id] = StatusResponse{downloadId.Id, start_time.String(), time.Now().String(), "SUCCESSFUL", "CONCURRENT", currentlyDownloadedFiles}
		return
	}()

	return downloadId
}

type DownloadId struct {
	Id string
}

var DownloadStatusMap = make(map[string]StatusResponse)
var UrlPathMap = make(map[string]string)
var currentlyDownloadedFiles = make(map[string]string)

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
	UrlPathMap[url] = filepath
	currentlyDownloadedFiles[url] = filepath
	return err
}

func GenerateUuid() string {
	id := uuid.New()
	return id.String()
}

func Worker(reqChan chan string, noOfUrls uint64, id string, counter *uint64) {
	for {
		select {
		case url, exists := <-reqChan:
			_ = DownloadFile(url)
			temp_status := DownloadStatusMap[id]
			temp_status.Files = currentlyDownloadedFiles
			DownloadStatusMap[id] = temp_status
			atomic.AddUint64(counter, 1)
			if !exists {
				return
			}
		}
		if atomic.LoadUint64(counter) == noOfUrls {
			close(reqChan)
			return
		}
	}
}
