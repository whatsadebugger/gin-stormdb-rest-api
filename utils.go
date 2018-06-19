package main

import (
	"io"
	"net/http"
)

func BasicPost(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}

// BasicGet uses the given username and password to send a GET request
// at the given URL and returns a response.
func BasicGet(url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(request)
	return resp, err
}

// BasicDelete sends a DELETE request to the HTTP client with the given username
// and password to enticate at the url with contentType and returns a response.
func BasicDelete(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}
