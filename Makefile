.PHONY: clean test security build run

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

bench:
	go test -bench=. -benchmem ./...

linter:
	golangci-lint run