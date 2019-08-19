package controller

import (
	"encoding/json"
	"github.com/dhairyapunjabi/FileDownloadManager/model"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func DownloadManager(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var downloadRequest model.Download
	json.Unmarshal(requestBody, &downloadRequest)
	if downloadRequest.Type == "serial" {
		for _, url := range downloadRequest.Urls {
			_ = DownloadFile(url)
		}
		downloadId := model.DownloadId{"Id" + generateUuid()}
		writer.Header().Set("Content-type", "application/json")
		id, _ := json.Marshal(downloadId)
		writer.Write(id)
	}
}

func DownloadFile(url string) error {
	filepath := "/tmp" + "/" + generateUuid()
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

func generateUuid() string {
	id := uuid.New()
	return id.String()
}
