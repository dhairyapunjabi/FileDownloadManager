package controller

import (
	"encoding/json"
	"github.com/dhairyapunjabi/FileDownloadManager/model"
	"io/ioutil"
	"net/http"
	"time"
)

type DownloadRequest struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

func DownloadManager(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var downloadRequest DownloadRequest
	json.Unmarshal(requestBody, &downloadRequest)
	if downloadRequest.Type == "serial" {
		serialDownload := model.SerialDownload{Urls: downloadRequest.Urls}
		start_time := time.Now()
		serialDownload.DownloadUrls()
		end_time := time.Now()
		downloadId := model.DownloadId{"Id" + model.GenerateUuid()}
		filesmap := make(map[string]string)
		for _, url := range serialDownload.Urls {
			filesmap[url] = model.UrlPathMap[url]
		}
		model.DownloadStatusMap[downloadId.Id] = model.StatusResponse{downloadId.Id, start_time.String(), end_time.String(), "SUCCESSFUL", "SERIAL", filesmap}
		writer.Header().Set("Content-type", "application/json")
		id, _ := json.Marshal(downloadId)
		writer.Write(id)
	}
}
