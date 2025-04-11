package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Start a goroutine to print a message every 10 seconds
	go func() {
		for {
			fmt.Println("Hello, this is a message printed every 10 seconds.")
			time.Sleep(10 * time.Second)
		}
	}()

	// Set up a simple HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the server!")
	})

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}