package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func GetHttpJsonResponse[T any](hc *HttpClient, url string, obj *T) error {
	// func GetHttpJsonResponse(hc *HttpClient, url string, obj *User) (*User, error) {

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

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("http response status code not 200. Status code: %v, Status: %s", resp.StatusCode, resp.Status)
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
