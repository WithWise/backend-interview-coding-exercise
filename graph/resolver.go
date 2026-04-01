package graph

// This file will not be regenerated automatically.
// It serves as dependency injection for the resolver.

import (
	"github.com/wise/backend-interview-kit/internal/repository"
	"github.com/wise/backend-interview-kit/internal/service"
)

// Resolver is the root GraphQL resolver. It holds shared dependencies that
// are injected into query, mutation, and field resolvers.
type Resolver struct {
	repo   repository.Repository
	jobSvc service.JobService
}

// NewResolver constructs a Resolver with the given dependencies.
func NewResolver(repo repository.Repository, jobSvc service.JobService) *Resolver {
	return &Resolver{repo: repo, jobSvc: jobSvc}
}
