package service

import (
	"github.com/wise/backend-interview-kit/internal/repository"
)

// JobService handles the business logic for job management.
//
// TODO: Define the methods you need on this interface and implement them
// in jobService below. Think about what belongs here vs. in the resolver.
type JobService interface {
	// TODO: add methods
}

type jobService struct {
	repo repository.Repository
}

// NewJobService returns a new JobService.
func NewJobService(repo repository.Repository) JobService {
	return &jobService{repo: repo}
}

// TODO: implement methods on jobService.
//
// Hints:
//   - AssignJob business rules: job must be PENDING; driver must not already
//     have an IN_PROGRESS job.
//   - UpdateJobStatus business rules: only the transitions defined in
//     model.ValidTransitions are permitted.
//   - Return descriptive errors — the resolver will surface these to the caller.

var _ JobService = (*jobService)(nil) // compile-time interface check
