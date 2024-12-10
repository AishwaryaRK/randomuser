package main

import (
	"fmt"
	"log"
)

type Response struct {
	Results []User `json:"results"`
}

type User struct {
	Name Name `json:"name"`

	Location Location `json:"location"`
	Email    string   `json:"email"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
}

type Location struct {
	City        string      `json:"city"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func GetRandomUser(hc *HttpClient, url string) (*User, error) {
	var resp Response
	err := GetHttpJsonResponse(hc, url, &resp)
	if err != nil {
		return nil, err
	}
	if len(resp.Results) == 0 {
		err := fmt.Errorf("empty http response user result")
		log.Println(err)
		return nil, err
	}

	return &resp.Results[0], nil
}
