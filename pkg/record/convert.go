package record

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

var (
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

		ifVal := rv.Field(i).Interface()
		if nested, ok := ifVal.(proto.Message); ok {
			convertedNested, err := Convert(nested)
			if err != nil {
				return nil, err
			}
			converted[name] = convertedNested
		} else if nestedArray, ok := ifVal.([]proto.Message); ok {
			arr := []map[string]interface{}{}
			for _, nested := range nestedArray {
				convertedNested, err := Convert(nested)
				if err != nil {
					return nil, err
				}
				arr = append(arr, convertedNested)
			}
			converted[name] = arr
		} else {
			converted[name] = ifVal
		}
	}

	return converted, nil
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
