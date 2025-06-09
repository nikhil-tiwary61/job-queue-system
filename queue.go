package main

import (
	"fmt"
	"math/rand"
	"time"
)

func StartDispatcher(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go worker(i)
	}

	go func() {
		for job := range jobQueue {
			updateJobStatus(job.ID, StatusRunning)

			time.Sleep(2 * time.Second)

			updateJobStatus(job.ID, StatusDone)
		}
	}()
}

func worker(id int) {
	fmt.Printf("Worker %d started\n", id)
	for job := range jobQueue {
		fmt.Printf("Worker %d picked job %s\n", id, job.ID)
		updateJobStatus(job.ID, StatusRunning)

		time.Sleep(2 * time.Second)

		if rand.Intn(5) == 0 {
			fmt.Printf("Worker %d: job %s failed\n", id, job.ID)
			updateJobStatus(job.ID, StatusFailed)
		} else {
			fmt.Printf("Worker %d: job %s done\n", id, job.ID)
			updateJobStatus(job.ID, StatusDone)
		}
	}
}

func updateJobStatus(jobID string, status JobStatus) {
	storeMutex.Lock()
	job := jobStore[jobID]
	job.Status = status
	jobStore[jobID] = job
	storeMutex.Unlock()
}
