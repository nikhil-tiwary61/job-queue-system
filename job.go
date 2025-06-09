package main

import (
	"time"

	"github.com/google/uuid"
)

type JobStatus string

const (
	StatusQueued  JobStatus = "QUEUED"
	StatusRunning JobStatus = "RUNNING"
	StatusDone    JobStatus = "DONE"
	StatusFailed  JobStatus = "FAILED"
)

type Job struct {
	ID        string    `json:"id"`
	Status    JobStatus `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Payload   string    `json:"payload,omitempty"`
}

func NewJob() Job {
	return Job{
		ID:        uuid.New().String(),
		Status:    StatusQueued,
		CreatedAt: time.Now(),
	}
}
