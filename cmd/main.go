package main

import (
	"github.com/dhairyapunjabi/FileDownloadManager/route"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	route.RouteRequest(server)
	_ = http.ListenAndServe(":8000", server)
}
