# avro-protobuf

![Go](https://github.com/syucream/avro-protobuf/workflows/Go/badge.svg)

A schema and value conversion module and CLI tools in Go, like [avro-protobuf](https://avro.apache.org/docs/1.8.2/api/java/org/apache/avro/protobuf/package-summary.html).

It bundles protobuf <-> avro schema conversions, record value conversions, SerDe and protoc plugin.

## Supported conversions

### schema

- protobuf descriptors -> Avro schema
  - supported
- Avro schema -> protobuf descriptors
  - not yet
  
### value

- protobuf value -> Avro bytes
  - partially
- Avro bytes -> protobuf value
  - not yet
- protobuf bytes -> Avro value
  - not yet
- Avro value -> protobuf bytes
  - partially
  
## Usage

### protobuf value <-> Avro bytes

```go
    msg := &your_proto_message{}

    serDe, err := serde.NewSerDe(msg)
    if err != nil {
        t.Error(err)
    }

    bin, err = serDe.Serialize(msg)
    if err != nil {
        t.Error(err)
    }
    // Got avro bytes!
    
    msg2 := &your_proto_message{}
    err = serDe.Deserialize(bin, msg2)
    if err != nil {
        t.Error(err)
    }
    // Got protobuf struct value!
```

### protobuf bytes <-> Avro value

- TODO

## protoc-gen-avro

It's a subproject of the avro-protobuf. It's a protoc plugin read .proto files and generate .avsc files.
You can get a binary of the plugin at the release page of this GitHub repo, or `go get`

```sh
$ go get -u github.com/syucream/avro-protobuf/tree/master/cmd/protoc-gen-avro
```

It works as a protoc plugin, like:

```sh
$ cat proto/com/syucream/example/simple.proto
// https://developers.google.com/protocol-buffers/docs/proto3#simple

syntax = "proto3";

message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
 
$ protoc --plugin=./protoc-gen-avro --avro_out=./gen proto/com/syucream/example/simple.proto

$ cat gen/SearchRequest.avsc
{
  "fields": [
    {
      "default": "",
      "name": "query",
      "type": "string"
    },
    {
      "default": 0,
      "name": "page_number",
      "type": "int"
    },
    {
      "default": 0,
      "name": "result_per_page",
      "type": "int"
    }
  ],
  "name": "SearchRequest",
  "namespace": "google.protobuf",
  "type": "record"
}
```

## references

- [goavro](https://github.com/linkedin/goavro)
