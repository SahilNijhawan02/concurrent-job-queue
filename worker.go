package main

import (
	"fmt"
	"sync"
	"time"
)

func startWorkers(n int, wg *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, wg)
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobQueue {
		fmt.Println("Worker", id, "processing job", job.ID)

		time.Sleep(10 * time.Second)

		job.Status = "completed"
		job.Result = "done"

		mu.Lock()
		jobStore[job.ID] = job
		mu.Unlock()
	}
}
