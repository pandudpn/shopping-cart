run:
	go run cmd/main.go
test:
	go test ./... --coverprofile coverage.out -v