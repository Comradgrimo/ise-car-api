.PHONY: build
build:
	go build cmd/ise-car-api/main.go

.PHONY: test
test:
	go test -v ./...