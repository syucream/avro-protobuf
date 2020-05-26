package record

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/syucream/avro-protobuf/pkg/schema"

	"github.com/golang/protobuf/descriptor"
	"github.com/linkedin/goavro"

	"github.com/golang/protobuf/proto"
)

const unionKeyPrefix = "__unionkey__"

var (
	unionKeyPrefixLen = len(unionKeyPrefix)

	ErrUnknownStructTagFormat = fmt.Errorf("struct tag format is broken")
)

// Convert gets proto message and converts to json like structure
func Convert(v proto.Message) (map[string]interface{}, error) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	rt := rv.Type()

	converted := map[string]interface{}{}
	for i := 0; i < rt.NumField(); i++ {
		// Skip implicit values
		fieldName := rt.Field(i).Name
		if len(fieldName) >= 4 && fieldName[:4] == "XXX_" {
			continue
		}

		name, err := getFieldNameFromTag(rt.Field(i).Tag)
		if err != nil {
			return nil, err
		}

		fv := rv.Field(i)

		// []interface{} is special because type assertion doesn't work
		if fv.Kind() == reflect.Slice {
			arr := []interface{}{}
			for j := 0; j < fv.Len(); j++ {
				if nested, ok := fv.Index(j).Interface().(descriptor.Message); ok {
					convertedNested, err := Convert(nested)
					if err != nil {
						return nil, err
					}
					ns := schema.GetNamespace(proto.MessageName(nested))
					_, md := descriptor.ForMessage(nested)
					unionKey := ns + "." + md.GetName()
					arr = append(arr, goavro.Union(unionKey, convertedNested))
				}
			}
			converted[name] = arr
		} else {
			ifVal := fv.Interface()
			if nested, ok := ifVal.(descriptor.Message); ok {
				convertedNested, err := Convert(nested)
				if err != nil {
					return nil, err
				}
				ns := schema.GetNamespace(proto.MessageName(nested))
				_, md := descriptor.ForMessage(nested)
				unionKey := ns + "." + md.GetName()
				converted[name] = goavro.Union(unionKey, convertedNested)
			} else {
				converted[name] = ifVal
			}
		}
	}

	return converted, nil
}

func IsUnionKey(k string) bool {
	return len(k) > unionKeyPrefixLen && k[:unionKeyPrefixLen] == unionKeyPrefix
}

func getFieldNameFromTag(tag reflect.StructTag) (string, error) {
	protoTag := tag.Get("protobuf")

	values := strings.Split(protoTag, ",")
	if len(values) < 4 {
		return "", ErrUnknownStructTagFormat
	}

	maybeName := strings.Split(values[3], "=")
	if len(values) < 2 {
		return "", ErrUnknownStructTagFormat
	}

	return maybeName[1], nil
}
