.PHONY: proto
proto:
	find proto -name "*.proto" -print0 | xargs -0 -IXXX protoc --go_out=gen/ XXX

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
