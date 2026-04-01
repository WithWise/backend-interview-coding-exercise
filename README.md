# Wise Backend Interview Exercise

A 45-minute live coding exercise for senior backend engineers.

## The scenario

You're building a **Job Management API** for Wise's logistics platform. Drivers pick up and deliver cargo between depots. The API is a GraphQL service written in Go.

The project is already scaffolded and compiles out of the box. Your job is to implement the business logic.

---

## Getting started

**Prerequisites:** Go 1.23+

```bash
go run .
```

The server starts on **http://localhost:8080** with a GraphQL playground where you can test your work interactively.

---

## What's already built

| File/Package | Description |
|---|---|
| `graph/schema.graphqls` | Full GraphQL schema |
| `graph/schema.resolvers.go` | Resolver stubs — **this is where you'll work** |
| `graph/resolver.go` | Dependency injection for the resolver |
| `internal/model/` | Domain models (`Job`, `Driver`, `Depot`, `JobStatus`) |
| `internal/repository/` | Repository interface + in-memory implementation (pre-seeded with data) |
| `internal/service/job_service.go` | Service skeleton — **you'll implement the business logic here** |

The `drivers` and `driver` queries are **pre-implemented** as working examples of the patterns used in this codebase.

---

## Tasks

See **[TASKS.md](./TASKS.md)** for the full exercise brief.

---

## Seed data

The in-memory repository is pre-seeded with:

- **3 depots:** Birmingham Central, Solihull Hub, Coventry Distribution Centre
- **4 drivers:** Tom McIntosh (DRV-1), Dominic Szabad (DRV-2), Tariq Avila (DRV-3), Dean Nicklin (DRV-4)
- **8 jobs** in various statuses — including `JOB-3` already `IN_PROGRESS` and assigned to Dominic Szabad

Useful IDs to test with: `JOB-1`, `JOB-3`, `DRV-1`, `DRV-2`

---

## Example query

Once the server is running, try this in the playground:

```graphql
{
  drivers {
    id
    name
    licenceNumber
  }
}
```
