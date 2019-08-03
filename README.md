# avro-protobuf

An [avro-protobuf](https://avro.apache.org/docs/1.8.2/api/java/org/apache/avro/protobuf/package-summary.html) inplementation in Go.

## Supported conversions

- protobuf value -> Avro bytes
  - partially
- Avro bytes -> protobuf value
  - no
- protobuf bytes -> Avro value
  - no
- Avro value -> protobuf bytes
  - no
  
## Usage

### protobuf value -> Avro bytes

```go
    msg := ptypes.TimestampNow()
    serializer, err := serde.NewSerDe(msg)
    if err != nil {
        t.Error(err)
    }

    bin, err = serializer.Serialize(msg)
    if err != nil {
        t.Error(err)
    }
    // Got avro bytes!
```

## protoc-gen-avro

It's a subproject of the avro-protobuf. It's a protoc plugin read .proto files and generate .avsc files.o

```sh
# when you have proto/nested.proto
$ protoc --plugin=./protoc-gen-avro --avro_out=./gen proto/nested.proto

# then you'll get .avsc in ./gen
$ ls gen/
SearchResponse.avsc
```

