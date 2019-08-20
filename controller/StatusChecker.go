package controller

import (
	"encoding/json"
	"github.com/dhairyapunjabi/FileDownloadManager/model"
	"net/http"
)

func StatusChecker(writer http.ResponseWriter, request *http.Request) {
	prefix := "/downloads/"
	id := request.URL.Path[len(prefix):]
	var statusResponse model.StatusResponse
	statusResponse, exists := model.DownloadStatusMap[id]
	if exists == true {
		writer.Header().Set("Content-type", "application/json")
		response, _ := json.Marshal(statusResponse)
		writer.Write(response)
	}
}
