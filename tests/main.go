package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const SERVER_URL = "http://localhost:8080"

var okUrls = []string{
	"/users",
	"/comments",
	"/posts",
}

var notFoundUrls = []string{
	"/users?test=trigger-not-found",
	"/comments?test=trigger-not-found",
	"/posts?test=trigger-not-found",
}

var forbiddenUrls = []string{
	"/users?test=trigger-forbidden",
	"/comments?test=trigger-forbidden",
	"/posts?test=trigger-forbidden",
}

var serverErrorUrls = []string{
	"/users?test=trigger-server-error",
	"/comments?test=trigger-server-error",
	"/posts?test=trigger-server-error",
}

func simulateRequests() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)
	for _, url := range okUrls {
		for range 10 {
			go func(url string) {
				for {
					fmt.Println("Sending request to", SERVER_URL+url)
					http.Get(SERVER_URL + url)
					time.Sleep(1 * time.Second)
				}
			}(url)
		}

	}
	for _, url := range notFoundUrls {
		for range 10 {
			go func(url string) {
				for {
					fmt.Println("Sending request to", SERVER_URL+url)
					http.Get(SERVER_URL + url)
					time.Sleep(2 * time.Second)
				}
			}(url)
		}
	}
	for _, url := range forbiddenUrls {
		for range 10 {
			go func(url string) {
				for {
					fmt.Println("Sending request to", SERVER_URL+url)
					http.Get(SERVER_URL + url)
					time.Sleep(3 * time.Second)
				}
			}(url)
		}
	}
	for _, url := range serverErrorUrls {
		for range 10 {
			go func(url string) {
				for {
					fmt.Println("Sending request to", SERVER_URL+url)
					http.Get(SERVER_URL + url)
					time.Sleep(4 * time.Second)
				}
			}(url)
		}
	}
	fmt.Println("Simulating requests...  Press Ctrl+C to stop.")

	<-stop
	fmt.Println("Shutting down...")
}

func main() {
	simulateRequests()
}
