package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/wise/backend-interview-kit/graph"
	"github.com/wise/backend-interview-kit/internal/repository"
	"github.com/wise/backend-interview-kit/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewMemoryRepository()
	jobSvc := service.NewJobService(repo)

	schema := graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(repo, jobSvc),
	})

	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("Wise Job Management API", "/query"))
	http.Handle("/query", srv)

	// TODO (gRPC bonus): start the gRPC server on port 9090.
	// Uncomment once you have generated the proto stubs and removed the build tag
	// from internal/grpc/server.go.
	//
	// lis, err := net.Listen("tcp", ":9090")
	// if err != nil {
	// 	log.Fatalf("failed to listen on :9090: %v", err)
	// }
	// grpcSrv := grpc.NewServer()
	// pb.RegisterDriverJobServiceServer(grpcSrv, grpcserver.NewServer(repo))
	// log.Println("gRPC server listening on :9090")
	// go grpcSrv.Serve(lis)

	log.Printf("Server running at http://localhost:%s/", port)
	log.Printf("GraphQL playground: http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
