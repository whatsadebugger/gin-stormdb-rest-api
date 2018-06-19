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

func BasicGet(url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(request)
	return resp, err
}

func BasicDelete(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}

func BasicPut(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("PUT", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}

func PostTextCSV(url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", "text/csv")

	resp, err := client.Do(request)
	return resp, err
}
