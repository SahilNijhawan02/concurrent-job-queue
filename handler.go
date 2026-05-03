package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func createJob(w http.ResponseWriter, r *http.Request) {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)

	job := Job{
		ID:     id,
		Status: "queued",
	}

	mu.Lock()
	jobStore[id] = job
	mu.Unlock()

	jobQueue <- job

	json.NewEncoder(w).Encode(job)
}

func getJob(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mu.Lock()
	job, ok := jobStore[id]
	mu.Unlock()

	if !ok {
		http.Error(w, "Job not found", 404)
		return
	}

	json.NewEncoder(w).Encode(job)
}
