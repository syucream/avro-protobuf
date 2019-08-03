.PHONY: proto
proto:
	protoc --go_out=gen/ proto/*.proto

.PHONY: fmt
fmt:
	goimports -w cmd/**/*.go
	goimports -w pkg/**/*.go
	gofmt -w cmd/**/*.go
	gofmt -w pkg/**/*.go

.PHONY: cmd
cmd:
	go build -o protoc-gen-avro cmd/protoc-gen-avro/main.go

.PHONY: test
test:
	go test -v ./...
