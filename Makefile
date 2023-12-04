run_all:
	go run cmd/adventofcode/main.go

run_one:
	go run cmd/runner/main.go

test:
	go test -v advent-of-code-2023/internal/challenges
