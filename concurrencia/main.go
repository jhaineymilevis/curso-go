package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println("Start: ", start)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.dropi.co",
		"https://www.instagram.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://api.github.com",
		"https://test.noexiste.com",
	}

	//declaring a chanel for comunicate go runtimes
	channel := make(chan string)

	for _, api := range apis {
		// go runtimes
		go checkApi(api, channel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-channel) //print channel data
	}

	//suspence time to see the result
	//time.Sleep(15 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Finished on : %v\n seconds", elapsed.Seconds())

}

func checkApi(api string, channel chan string) {

	_, err := http.Get(api)

	if err != nil {
		//fmt.Printf("Error Api Down %s: %v\n", api, err)             // prit bny console
		channel <- fmt.Sprintf("ERROR! Api Down %s: %v\n", api, err) //send data to chanel

		return
	}

	//fmt.Printf("Api esta en funcionamiento %s: %v\n", api, err)             // prit bny console
	channel <- fmt.Sprintf("SUCCESS: Api esta en funcionamiento %s: %v\n", api, err) //send data to chanel

}
