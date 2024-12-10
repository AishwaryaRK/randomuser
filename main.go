package main

import (
	"fmt"
	"log"
)

func main() {
	url := "https://randomuser.me/api/"
	timeout := 5
	hc := NewHttpClient(timeout)
	user, err := GetRandomUser(hc, url)
	if err != nil {
		log.Printf("Error getting random user")
		return
	}
	fmt.Println("User: ", user)
}
