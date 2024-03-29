package model

type StatusResponse struct {
	ID           string            `json:"id"`
	StartTime    string            `json:"start_time"`
	EndTime      string            `json:"end_time"`
	Status       string            `json:"status"`
	DownloadType string            `json:"download_type"`
	Files        map[string]string `json:"files"`
}
