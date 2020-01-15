package schema

import (
	"strings"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	genDescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type typeMap map[string]*genDescriptor.DescriptorProto

func GetRecordSchemaFromMessage(msg descriptor.Message) (map[string]interface{}, error) {
	fd, md := descriptor.ForMessage(msg)

	ns := GetNamespace(proto.MessageName(msg))

	return GetRecordSchemaFromDescriptor(fd, md, ns)
}

func GetRecordSchemaFromDescriptor(fd *genDescriptor.FileDescriptorProto, md *genDescriptor.DescriptorProto, ns string) (map[string]interface{}, error) {
	var fields []map[string]interface{}
	for _, f := range md.Field {
		rf, err := GetRecordSchemaFieldFromDescriptor(f, fd, getTypeMap(fd, md))
		if err != nil {
			return nil, err
		}
		fields = append(fields, rf)
	}

	return map[string]interface{}{
		"name": md.GetName(),
		// should it use fully-qualified?
		"namespace": ns,
		"type":      "record",
		"fields":    fields,
	}, nil
}

func GetNestedRecordSchemaFromDescriptor(fd *genDescriptor.FileDescriptorProto, md *genDescriptor.DescriptorProto, ns string) ([]interface{}, error) {
	recordSchema, err := GetRecordSchemaFromDescriptor(fd, md, ns)
	if err != nil {
		return nil, err
	}

	return []interface{}{"null", recordSchema}, nil
}

func GetRecordSchemaFieldFromDescriptor(d *genDescriptor.FieldDescriptorProto, fd *genDescriptor.FileDescriptorProto, tm typeMap) (map[string]interface{}, error) {
	var tpe interface{}
	if d.GetType() == genDescriptor.FieldDescriptorProto_TYPE_MESSAGE {
		// nested messages -> nested records
		tn := d.GetTypeName()
		if len(tn) == 0 || tn[0] != '.' {
			return nil, ErrUnknownProtoType
		}
		md, typeOk := tm[tn]
		if !typeOk {
			return nil, ErrUnknownProtoType
		}
		ns := GetNamespace(tn)
		t, err := GetNestedRecordSchemaFromDescriptor(fd, md, ns)
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

	if d.GetLabel() == genDescriptor.FieldDescriptorProto_LABEL_REPEATED {
		// repeated -> array
		return map[string]interface{}{
			"name":    d.GetName(),
			"type":    wrapByArray(tpe),
			"default": []interface{}{},
		}, nil
	} else {
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

	// TODO resolve messages recursively

	return messageTypes
}

// GetNamespace gets 'namespace' field in Avro from proto message name.
func GetNamespace(fullName string) string {
	sp := strings.Split(fullName, ".")

	// strip heading '.'
	if sp[0] == "" {
		sp = sp[1:]
	}

	return strings.Join(sp[:len(sp)-1], ".")
}

func wrapByArray(items interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type":  "array",
		"items": items,
	}
}
