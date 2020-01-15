# avro-protobuf

A schema and value conversion utilities in Go, like [avro-protobuf](https://avro.apache.org/docs/1.8.2/api/java/org/apache/avro/protobuf/package-summary.html).

It bundles protobuf <-> avro schema conversions, record value conversions, SerDe and protoc plugin.

## Supported conversions

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

```sh
# when you have proto/nested.proto
$ protoc --plugin=./protoc-gen-avro --avro_out=./gen proto/com/syucream/example/nested.proto

# then you'll get .avsc in ./gen
$ ls gen/
SearchResponse.avsc
$ cat gen/SearchResponse.avsc
{
  "fields": [
    {
      "default": [],
      "name": "results",
      "type": {
        "items": [
          "null",
          {
            "fields": [
              {
                "default": "",
                "name": "url",
                "type": "string"
              },
              {
                "default": "",
                "name": "title",
                "type": "string"
              },
              ...

```

## references

- [goavro](https://github.com/linkedin/goavro)
