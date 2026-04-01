package model

import "time"

type JobStatus string

const (
	JobStatusPending    JobStatus = "PENDING"
	JobStatusInProgress JobStatus = "IN_PROGRESS"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
)

// ValidTransitions defines the allowed job status state machine.
// PENDING → IN_PROGRESS → COMPLETED
//
//	              ↘ FAILED
var ValidTransitions = map[JobStatus][]JobStatus{
	JobStatusPending:    {JobStatusInProgress},
	JobStatusInProgress: {JobStatusCompleted, JobStatusFailed},
}

type Job struct {
	ID            string
	Reference     string
	Status        JobStatus
	OriginID      string
	DestinationID string
	DriverID      *string
	CreatedAt     time.Time
}

type Driver struct {
	ID            string
	Name          string
	LicenceNumber string
}

type Depot struct {
	ID      string
	Name    string
	Address string
}

type JobsFilter struct {
	Status   *JobStatus
	DriverID *string
}
