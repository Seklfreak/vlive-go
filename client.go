package vlive_go

import (
	"net/http"
)

type VLive struct {
	httpClient *http.Client
}

func NewVLive(httpClient *http.Client) *VLive {
	return &VLive{
		httpClient: httpClient,
	}
}
