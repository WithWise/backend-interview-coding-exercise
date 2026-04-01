package repository

import "github.com/wise/backend-interview-kit/internal/model"

// Repository defines the data access interface for the job management service.
type Repository interface {
	GetJob(id string) (*model.Job, error)
	ListJobs(filter model.JobsFilter) ([]*model.Job, error)
	UpdateJob(job *model.Job) (*model.Job, error)

	GetDriver(id string) (*model.Driver, error)
	ListDrivers() ([]*model.Driver, error)

	GetDepot(id string) (*model.Depot, error)
}
