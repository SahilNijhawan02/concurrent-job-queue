package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	startWorkers(20, &wg)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/jobs", createJob)
	http.HandleFunc("/job", getJob)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Println("Server running on port:", port)

	http.ListenAndServe(":"+port, nil)

	// Graceful shutdown (not triggered here yet)
	close(jobQueue)
	wg.Wait()
}
