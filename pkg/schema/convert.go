package schema

import (
	"github.com/golang/protobuf/descriptor"
	genDescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type typeMap map[string]*genDescriptor.DescriptorProto

func GetRecordSchemaFromMessage(msg descriptor.Message) (map[string]interface{}, error) {
	fd, md := descriptor.ForMessage(msg)

	return GetRecordSchemaFromDescriptor(fd, md)
}

func GetRecordSchemaFromDescriptor(fd *genDescriptor.FileDescriptorProto, md *genDescriptor.DescriptorProto) (map[string]interface{}, error) {
	var fields []map[string]interface{}
	for _, f := range md.Field {
		rf, err := GetRecordSchemaFieldFromDescriptor(f, fd, getTypeMap(fd, md))
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

func GetNestedRecordSchemaFromDescriptor(fd *genDescriptor.FileDescriptorProto, md *genDescriptor.DescriptorProto) ([]interface{}, error) {
	recordSchema, err := GetRecordSchemaFromDescriptor(fd, md)
	if err != nil {
		return nil, err
	}

	return []interface{}{"null", recordSchema}, nil
}

// TODO support repeated
func GetRecordSchemaFieldFromDescriptor(d *genDescriptor.FieldDescriptorProto, fd *genDescriptor.FileDescriptorProto, tm typeMap) (map[string]interface{}, error) {
	var tpe interface{}
	if d.GetType() == genDescriptor.FieldDescriptorProto_TYPE_MESSAGE {
		tn := d.GetTypeName()
		if len(tn) == 0 || tn[0] != '.' {
			return nil, ErrUnknownProtoType
		}
		md, typeOk := tm[tn]
		if !typeOk {
			return nil, ErrUnknownProtoType
		}
		t, err := GetNestedRecordSchemaFromDescriptor(fd, md)
		if err != nil {
			return nil, err
		}
		tpe = t
	} else {
		t, typeOk := ProtoType2AvroType[d.GetType()]
		if !typeOk {
			return nil, ErrUnknownProtoType
		}
		tpe = t
	}

	dv, defaultOk := ProtoType2AvroDefault[d.GetType()]
	if !defaultOk {
		return nil, ErrUnknownProtoType
	}

	return map[string]interface{}{
		"name":    d.GetName(),
		"type":    tpe,
		"default": dv,
	}, nil
}

func getTypeMap(fd *genDescriptor.FileDescriptorProto, md *genDescriptor.DescriptorProto) typeMap {
	messageTypes := typeMap{}

	// file scope
	for _, mt := range fd.GetMessageType() {
		fqtn := "." + fd.GetPackage() + "." + mt.GetName()
		messageTypes[fqtn] = mt
	}

	// nested message scope
	for _, nmt := range md.GetNestedType() {
		fqtn := "." + fd.GetPackage() + "." + md.GetName() + "." + nmt.GetName()
		messageTypes[fqtn] = nmt
	}

	// TODO it doesn't resolve messages recursively for now

	return messageTypes
}
