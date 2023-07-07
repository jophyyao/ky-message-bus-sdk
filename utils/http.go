package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpPostJson(url string, jsonStr []byte) (string, error) {
	// jsonStr :=[]byte(`{ "username": "auto", "password": "auto123123" }`)
	// url:= "https://www.denlery.top/api/v1/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return "", err
	}
	defer resp.Body.Close()

	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), err
}
