# aninmals-race

## Prerequisities

- Docker
- Go 1.21

## Run

1. `docker run -p 6382:6379 --name aninmals-race -d redis`
2. `go run cmd/race/main.go`