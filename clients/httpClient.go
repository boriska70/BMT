package clients

import (
	"net/http"
	"time"
)

func CreateHttpClient() *http.Client {

	return &http.Client{
		Transport: setTransport(),
	}
}

func setTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}
}
