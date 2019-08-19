package route

import (
	"fmt"
	"net/http"
)

func RouteRequest(server *http.ServeMux) {
	//handle /health api
	server.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "OK")
	})
}
