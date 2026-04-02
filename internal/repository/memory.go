package repository

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/wise/backend-interview-kit/internal/model"
)

var ErrNotFound = errors.New("item not found")

type memoryRepository struct {
	mu      sync.RWMutex
	jobs    map[string]*model.Job
	drivers map[string]*model.Driver
	depots  map[string]*model.Depot
}

// NewMemoryRepository returns a Repository backed by in-memory data,
// pre-seeded with depots, drivers, and jobs for the exercise.
func NewMemoryRepository() Repository {
	r := &memoryRepository{
		jobs:    make(map[string]*model.Job),
		drivers: make(map[string]*model.Driver),
		depots:  make(map[string]*model.Depot),
	}
	r.seed()
	return r
}

func (r *memoryRepository) seed() {
	// Depots
	depots := []*model.Depot{
		{ID: "DEPOT-1", Name: "Birmingham Central", Address: "Saltley Industrial Estate, Birmingham B8 1AA"},
		{ID: "DEPOT-2", Name: "Solihull Hub", Address: "Lode Lane, Solihull B91 2LQ"},
		{ID: "DEPOT-3", Name: "Coventry Distribution Centre", Address: "Middlemarch Business Park, Coventry CV3 4FJ"},
	}
	for _, d := range depots {
		r.depots[d.ID] = d
	}

	// Drivers
	driverID2 := "DRV-2"
	driverID1 := "DRV-1"
	driverID3 := "DRV-3"
	drivers := []*model.Driver{
		{ID: "DRV-1", Name: "Tom McIntosh", LicenceNumber: "MCI123456"},
		{ID: "DRV-2", Name: "Dominic Szabad", LicenceNumber: "SZA987654"},
		{ID: "DRV-3", Name: "Tariq Avila", LicenceNumber: "AVI555123"},
		{ID: "DRV-4", Name: "Dean Nicklin", LicenceNumber: "NIC111222"},
	}
	for _, d := range drivers {
		r.drivers[d.ID] = d
	}

	// Jobs
	base := time.Now().Add(-72 * time.Hour)
	jobs := []*model.Job{
		{
			ID:            "JOB-1",
			Reference:     "REF-2024-001",
			Status:        model.JobStatusPending,
			OriginID:      "DEPOT-1",
			DestinationID: "DEPOT-2",
			CreatedAt:     base.Add(0 * time.Hour),
		},
		{
			ID:            "JOB-2",
			Reference:     "REF-2024-002",
			Status:        model.JobStatusPending,
			OriginID:      "DEPOT-2",
			DestinationID: "DEPOT-3",
			CreatedAt:     base.Add(1 * time.Hour),
		},
		{
			ID:            "JOB-3",
			Reference:     "REF-2024-003",
			Status:        model.JobStatusInProgress,
			OriginID:      "DEPOT-1",
			DestinationID: "DEPOT-3",
			DriverID:      &driverID2,
			CreatedAt:     base.Add(2 * time.Hour),
		},
		{
			ID:            "JOB-4",
			Reference:     "REF-2024-004",
			Status:        model.JobStatusCompleted,
			OriginID:      "DEPOT-2",
			DestinationID: "DEPOT-1",
			DriverID:      &driverID1,
			CreatedAt:     base.Add(3 * time.Hour),
		},
		{
			ID:            "JOB-5",
			Reference:     "REF-2024-005",
			Status:        model.JobStatusPending,
			OriginID:      "DEPOT-3",
			DestinationID: "DEPOT-1",
			CreatedAt:     base.Add(4 * time.Hour),
		},
		{
			ID:            "JOB-6",
			Reference:     "REF-2024-006",
			Status:        model.JobStatusFailed,
			OriginID:      "DEPOT-1",
			DestinationID: "DEPOT-2",
			DriverID:      &driverID3,
			CreatedAt:     base.Add(5 * time.Hour),
		},
		{
			ID:            "JOB-7",
			Reference:     "REF-2024-007",
			Status:        model.JobStatusPending,
			OriginID:      "DEPOT-2",
			DestinationID: "DEPOT-3",
			CreatedAt:     base.Add(6 * time.Hour),
		},
		{
			ID:            "JOB-8",
			Reference:     "REF-2024-008",
			Status:        model.JobStatusPending,
			OriginID:      "DEPOT-3",
			DestinationID: "DEPOT-2",
			CreatedAt:     base.Add(7 * time.Hour),
		},
	}
	for _, j := range jobs {
		r.jobs[j.ID] = j
	}
}

func (r *memoryRepository) GetJob(id string) (*model.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	job, ok := r.jobs[id]
	if !ok {
		return nil, ErrNotFound
	}
	return job, nil
}

func (r *memoryRepository) ListJobs(filter model.JobsFilter) ([]*model.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var jobs []*model.Job
	for _, j := range r.jobs {
		if filter.Status != nil && j.Status != *filter.Status {
			continue
		}
		if filter.DriverID != nil {
			if j.DriverID == nil || *j.DriverID != *filter.DriverID {
				continue
			}
		}
		jobs = append(jobs, j)
	}
	return jobs, nil
}

func (r *memoryRepository) UpdateJob(job *model.Job) (*model.Job, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.jobs[job.ID]; !ok {
		return nil, fmt.Errorf("job %s not found", job.ID)
	}
	r.jobs[job.ID] = job
	return job, nil
}

func (r *memoryRepository) GetDriver(id string) (*model.Driver, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	driver, ok := r.drivers[id]
	if !ok {
		return nil, ErrNotFound
	}
	return driver, nil
}

func (r *memoryRepository) ListDrivers() ([]*model.Driver, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	drivers := make([]*model.Driver, 0, len(r.drivers))
	for _, d := range r.drivers {
		drivers = append(drivers, d)
	}
	return drivers, nil
}

func (r *memoryRepository) GetDepot(id string) (*model.Depot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	depot, ok := r.depots[id]
	if !ok {
		return nil, ErrNotFound
	}
	return depot, nil
}
