package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string, data interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = checkStatusCode(res)
	if err != nil {
		return err
	}

	return parseJsonResponse(res, data)
}

func checkStatusCode(res *http.Response) error {
	switch res.StatusCode / 100 {
	case 2:
		return nil
	default:
		return fmt.Errorf("HTTP request failed")
	}
}

func parseJsonResponse(res *http.Response, data interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}
