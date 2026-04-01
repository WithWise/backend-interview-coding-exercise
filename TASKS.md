# Exercise Tasks

You have **45 minutes**. Work through the tasks in order — each one builds on the previous. It's fine not to finish everything; we care more about how you approach each problem than how far you get.

Before you start, spend a minute reading through the existing code — particularly the pre-implemented `Drivers` / `Driver` resolvers in `graph/schema.resolvers.go` and the repository interface in `internal/repository/repository.go`.

---

## Task 1 — Query jobs (≈15 min)

Implement the `jobs` and `job` queries.

```graphql
query {
  jobs(filter: { status: PENDING }) {
    id
    reference
    status
    origin { name }
    destination { name }
    driver { name }
  }
}

query {
  job(id: "JOB-1") {
    id
    reference
    status
  }
}
```

**Requirements:**
- `jobs` returns all jobs when no filter is provided
- `jobs` supports optional filtering by `status` and/or `driverId`
- `job` returns `null` when the ID does not exist

**Where to work:**
- `graph/schema.resolvers.go` — `Jobs` and `Job` resolver methods
- `internal/service/job_service.go` — add the methods you need

---

## Task 2 — Assign a job (≈15 min)

Implement the `assignJob` mutation.

```graphql
mutation {
  assignJob(input: { jobId: "JOB-1", driverId: "DRV-1" }) {
    job {
      id
      status
      driver { name }
    }
    userError
  }
}
```

**Business rules:**
- The job must exist and be in `PENDING` status
- The driver must exist
- The driver must not already have an `IN_PROGRESS` job

On a rule violation, return a `userError` in the payload rather than a top-level GraphQL error.

**Useful test cases:**
- Assigning `JOB-1` to `DRV-1` should succeed
- Assigning `JOB-3` to `DRV-1` should fail (job is already `IN_PROGRESS`)
- Assigning `JOB-1` to `DRV-2` should fail (Dominic Szabad already has `JOB-3` in progress)

---

## Task 3 — Update job status (≈10 min)

Implement the `updateJobStatus` mutation.

```graphql
mutation {
  updateJobStatus(input: { jobId: "JOB-1", status: IN_PROGRESS }) {
    job { id status }
    userError
  }
}
```

**Valid transitions:**

```
PENDING → IN_PROGRESS → COMPLETED
                      ↘ FAILED
```

All other transitions are invalid and should return a `userError`. The valid transitions are also defined in `model.ValidTransitions` for reference.

---

## Task 4 — API design (≈10 min)

The product team has a new requirement:

> *"When viewing a driver's profile, we want to show their full job history."*

**Discuss** how you would extend the GraphQL schema to support this, then implement it if time allows.

Think about:
- Where does this belong in the schema? (New field on `Driver`? New top-level query?)
- What filtering or pagination might be needed?
- Are there any changes needed to the existing types?

---

## Bonus

If you finish early, pick any of the following:

- Add pagination to the `jobs` query (offset/limit or cursor-based — discuss the trade-offs)
- The schema currently has no way to distinguish between "not found" and "server error" in mutations — how would you improve the error model?

---

## Bonus — gRPC endpoint *(skip if unfamiliar with gRPC)*

> **Not familiar with gRPC? Skip this — it won't be held against you.**

Driver mobile apps need a lightweight way to fetch their currently active job. Expose this as a gRPC service running alongside the existing HTTP server.

**1. Define the proto**

`proto/jobs.proto` is pre-written for you:

```proto
syntax = "proto3";

package jobs;
option go_package = "github.com/wise/backend-interview-kit/proto/gen";

service DriverJobService {
  rpc GetActiveJob(GetActiveJobRequest) returns (GetActiveJobResponse);
}

message GetActiveJobRequest {
  string driver_id = 1;
}

message GetActiveJobResponse {
  string job_id     = 1;
  string reference  = 2;
  string status     = 3;
  string origin     = 4;
  string destination = 5;
}
```

`GetActiveJob` should return the driver's `IN_PROGRESS` job, or an empty response if they have none.

**2. Implement the server**

Create `internal/grpc/server.go` implementing the `DriverJobService` interface. Wire it up against the existing `repository.Repository` — no need to go through the service layer.

**3. Start the gRPC server**

In `main.go`, start a `grpc.NewServer()` on port `9090` in a goroutine alongside the existing HTTP server.

**Where to work:**
- `proto/jobs.proto` — proto definition is pre-written; generate Go stubs with:
  ```
  # Install the Go protoc plugins (once, if not already installed)
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

  # Generate stubs into proto/gen/
  protoc -I proto \
    --go_out=proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
    proto/jobs.proto

  # Pull the runtime dependency
  go get google.golang.org/grpc
  ```
- `internal/grpc/server.go` — skeleton is pre-written; remove the `//go:build ignore` tag at the top once stubs are generated, then fill in the `GetActiveJob` method
- `main.go` — uncomment the gRPC startup block near the bottom
