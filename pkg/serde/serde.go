package serde

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/linkedin/goavro"
	"github.com/syucream/avro-protobuf/pkg/record"
	"github.com/syucream/avro-protobuf/pkg/schema"
)

type SerDe struct {
	Codec       *goavro.Codec
	unmarshaler jsonpb.Unmarshaler
}

func NewSerDe(msg descriptor.Message) (*SerDe, error) {
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

	return &SerDe{
		Codec:       codec,
		unmarshaler: jsonpb.Unmarshaler{},
	}, nil
}

func (s *SerDe) Serialize(msg proto.Message) ([]byte, error) {
	recordMap, err := record.Convert(msg)
	if err != nil {
		return nil, err
	}

	return s.Codec.BinaryFromNative(nil, recordMap)
}

func (s *SerDe) Deserialize(avroBytes []byte, v proto.Message) error {
	datum, _, err := s.Codec.NativeFromBinary(avroBytes)
	if err != nil {
		return err
	}

	datumJson, err := json.Marshal(datum)
	if err != nil {
		return err
	}

	rbuf := bytes.NewBuffer(datumJson)

	return s.unmarshaler.Unmarshal(rbuf, v)
}
