.PHONY: proto
proto:
	protoc --go_out=gen/ proto/*.proto

.PHONY: fmt
fmt:
	gofmt -w pkg/**/*.go

.PHONY: test
test:
	go test -v ./...
