//go:build ignore
// Remove the build tag above once you have generated the proto stubs:
//
//   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//   protoc -I proto \
//     --go_out=proto/gen --go_opt=paths=source_relative \
//     --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
//     proto/jobs.proto
//   go get google.golang.org/grpc

package grpcserver

import (
	"context"
	"fmt"

	"github.com/wise/backend-interview-kit/internal/repository"
	pb "github.com/wise/backend-interview-kit/proto/gen"
)

// Server implements pb.DriverJobServiceServer.
type Server struct {
	pb.UnimplementedDriverJobServiceServer
	repo repository.Repository
}

// NewServer creates a new gRPC Server backed by the given repository.
func NewServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

// GetActiveJob returns the IN_PROGRESS job for the given driver, or an empty
// response if the driver has no active job.
func (s *Server) GetActiveJob(ctx context.Context, req *pb.GetActiveJobRequest) (*pb.GetActiveJobResponse, error) {
	// TODO: implement
	//
	// Suggested steps:
	//   1. Call s.repo.ListJobs with a filter on req.DriverId
	//   2. Find the job whose Status == model.JobStatusInProgress
	//   3. Fetch the origin and destination depots via s.repo.GetDepot
	//   4. Return a populated GetActiveJobResponse
	//
	// If the driver has no active job, return &pb.GetActiveJobResponse{}, nil
	return nil, fmt.Errorf("not implemented")
}
