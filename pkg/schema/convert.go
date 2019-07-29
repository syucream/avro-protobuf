package schema

import (
	"github.com/golang/protobuf/descriptor"
	genDescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func GetRecordSchemaFromMessage(msg descriptor.Message) (map[string]interface{}, error) {
	fd, md := descriptor.ForMessage(msg)

	return GetRecordSchemaFromDescriptor(*fd, *md)
}

func GetRecordSchemaFromDescriptor(fd genDescriptor.FileDescriptorProto, md genDescriptor.DescriptorProto) (map[string]interface{}, error) {
	var fields []map[string]interface{}
	for _, f := range md.Field {
		rf, err := GetRecordSchemaFieldFromDescriptor(*f)
		if err != nil {
			return nil, err
		}
		fields = append(fields, rf)
	}

	return map[string]interface{}{
		"name":      md.GetName(),
		"namespace": fd.GetPackage(),
		"type":      "record",
		"fields":    fields,
	}, nil
}

func GetRecordSchemaFieldFromDescriptor(d genDescriptor.FieldDescriptorProto) (map[string]interface{}, error) {
	t, typeOk := ProtoType2AvroType[d.GetType()]
	if !typeOk {
		return nil, ErrUnknownProtoType
	}

	dv, defaultOk := ProtoType2AvroDefault[d.GetType()]
	if !defaultOk {
		return nil, ErrUnknownProtoType
	}

	return map[string]interface{}{
		"name":    d.GetName(),
		"type":    t,
		"default": dv,
	}, nil
}
