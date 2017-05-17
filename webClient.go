package main

import (
	"net/http"
	"time"
)

var httpClient *http.Client

func checkHTTPStatus(URL string) bool {

	if nil == httpClient {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	res, err := httpClient.Get(URL)

	if nil != err {
		return false
	}

	if res.StatusCode != http.StatusOK {
		return false
	}

	return true
}
