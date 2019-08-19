package model

type Download struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type DownloadId struct {
	Id string
}
