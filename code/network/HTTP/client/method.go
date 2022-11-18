package main

import (
	"bytes"
	"io"
	"net/http"
)

func Req(method string,
	baseURL string, url string,
	data []byte,
	setters ...func(*http.Request)) (
	body []byte, header http.Header, statusCode int) {

	fullUrl := baseURL + url
	request, err := http.NewRequest(method, fullUrl, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	if method == "POST" || method == "PUT" || method == "DELETE" {
		request.Header.Set("Content-Type", "application/json;charset=utf8;")
		// request.Header.Set("Content-Type", "multipart/form-data;charset=utf8;")
	}

	return reqInner(request, setters...)
}

func reqInner(request *http.Request,
	setters ...func(*http.Request)) (
	body []byte, header http.Header, statusCode int) {

	for _, setter := range setters {
		setter(request)
	}

	res, err := HttpClient.Do(request)
	if err != nil {
		return
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	header = res.Header
	statusCode = res.StatusCode
	return
}
