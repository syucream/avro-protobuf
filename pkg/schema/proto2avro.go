package schema

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

var (
	// mapping proto type to avro primitive type
	ProtoType2AvroType = map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:  "double",
		descriptor.FieldDescriptorProto_TYPE_FLOAT:   "float",
		descriptor.FieldDescriptorProto_TYPE_INT64:   "long",
		descriptor.FieldDescriptorProto_TYPE_UINT64:  "long",
		descriptor.FieldDescriptorProto_TYPE_INT32:   "int",
		descriptor.FieldDescriptorProto_TYPE_FIXED64: "long",
		descriptor.FieldDescriptorProto_TYPE_FIXED32: "int",
		descriptor.FieldDescriptorProto_TYPE_BOOL:    "boolean",
		descriptor.FieldDescriptorProto_TYPE_STRING:  "string",
		// nested's
		// descriptor.FieldDescriptorProto_TYPE_GROUP:
		// descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		descriptor.FieldDescriptorProto_TYPE_BYTES:  "bytes",
		descriptor.FieldDescriptorProto_TYPE_UINT32: "int",
		// nested's
		// descriptor.FieldDescriptorProto_TYPE_ENUM:
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: "int",
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: "long",
		descriptor.FieldDescriptorProto_TYPE_SINT32:   "int",
		descriptor.FieldDescriptorProto_TYPE_SINT64:   "long",
	}

	// mapping proto type to avro zero-value to use it as default
	ProtoType2AvroDefault = map[descriptor.FieldDescriptorProto_Type]interface{}{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:  0,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:   0,
		descriptor.FieldDescriptorProto_TYPE_INT64:   0,
		descriptor.FieldDescriptorProto_TYPE_UINT64:  0,
		descriptor.FieldDescriptorProto_TYPE_INT32:   0,
		descriptor.FieldDescriptorProto_TYPE_FIXED64: 0,
		descriptor.FieldDescriptorProto_TYPE_FIXED32: 0,
		descriptor.FieldDescriptorProto_TYPE_BOOL:    false,
		descriptor.FieldDescriptorProto_TYPE_STRING:  "",
		// nested's
		// descriptor.FieldDescriptorProto_TYPE_GROUP:
		descriptor.FieldDescriptorProto_TYPE_MESSAGE: nil,
		descriptor.FieldDescriptorProto_TYPE_BYTES:  []byte{},
		descriptor.FieldDescriptorProto_TYPE_UINT32: 0,
		// nested's
		// descriptor.FieldDescriptorProto_TYPE_ENUM:
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: 0,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: 0,
		descriptor.FieldDescriptorProto_TYPE_SINT32:   0,
		descriptor.FieldDescriptorProto_TYPE_SINT64:   0,
	}
)
