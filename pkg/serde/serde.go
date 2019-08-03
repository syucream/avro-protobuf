package serde

import (
	"encoding/json"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/linkedin/goavro"
	"github.com/syucream/avro-protobuf/pkg/record"
	"github.com/syucream/avro-protobuf/pkg/schema"
)

type Serializer struct {
	Codec *goavro.Codec
}

func NewSerializer(msg descriptor.Message) (*Serializer, error) {
	schemaMap, err := schema.GetRecordSchemaFromMessage(msg)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(schemaMap)
	if err != nil {
		return nil, err
	}

	codec, err := goavro.NewCodec(string(data))
	if err != nil {
		return nil, err
	}

	return &Serializer{
		Codec: codec,
	}, nil
}

func (s *Serializer) Serialize(msg proto.Message) ([]byte, error) {
	recordMap, err := record.Convert(msg)
	if err != nil {
		return nil, err
	}

	return s.Codec.BinaryFromNative(nil, recordMap)
}
