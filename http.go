package main

import (
	"net/http"
)

var agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36"

func httpDoSomething(Method string, URL string) (*http.Response, error) {

	c := &http.Client{}

	request, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", agent)

	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func httpGet(URL string) (*http.Response, error) {
	return httpDoSomething("GET", URL)
}
