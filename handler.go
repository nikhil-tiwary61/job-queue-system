package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	jobQueue   = make(chan Job, 100)
	jobStore   = make(map[string]Job)
	storeMutex = sync.RWMutex{}
)

type JobRequest struct {
	Payload string `json:"payload"`
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req JobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	job := NewJob()
	job.Payload = req.Payload

	storeMutex.Lock()
	jobStore[job.ID] = job
	storeMutex.Unlock()

	jobQueue <- job

	log.Printf("New job submitted: %s\n", job.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"job_id": job.ID})
}

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID := vars["id"]

	storeMutex.RLock()
	job, exists := jobStore[jobID]
	storeMutex.RUnlock()

	if !exists {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
