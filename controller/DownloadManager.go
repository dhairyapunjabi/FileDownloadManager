package controller

import (
	"encoding/json"
	"github.com/dhairyapunjabi/FileDownloadManager/model"
	"io/ioutil"
	"net/http"
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
		serialDownload.DownloadUrls()
		downloadId := model.DownloadId{"Id" + model.GenerateUuid()}
		writer.Header().Set("Content-type", "application/json")
		id, _ := json.Marshal(downloadId)
		writer.Write(id)
	}
}
