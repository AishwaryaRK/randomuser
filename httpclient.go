package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient(timeout int) *HttpClient {
	return &HttpClient{client: &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	},
	}
}

func isValidUrl(u string) bool {
	urlOb, err := url.Parse(u)
	if err != nil {
		return false
	}
	if urlOb.Scheme == "" || urlOb.Host == "" {
		return false
	}
	return true
}

func GetHttpJsonResponse[T any](hc *HttpClient, url string, obj *T) error {
	if !isValidUrl(url) {
		err := fmt.Errorf("invalid url: %s", url)
		log.Printf("%v", err)
		return err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := fmt.Errorf("error creating new http request. Error: %v", err)
		log.Printf("%v", err)
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := hc.client.Do(req)
	if err != nil {
		err := fmt.Errorf("error fetching http response. Error: %v", err)
		log.Printf("%v", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err := fmt.Errorf("http response status code not 2xx. Status code: %d, Status: %s", resp.StatusCode, resp.Status)
		log.Printf("%v", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("error reading http response body. Error: %v", err)
		log.Printf("%v", err)
		return err
	}

	// fmt.Println("----- ", string(body))
	err = json.Unmarshal(body, obj)
	if err != nil {
		err := fmt.Errorf("error unmarshaling json response. Error: %v", err)
		log.Printf("%v", err)
		return err
	}
	return nil
}
