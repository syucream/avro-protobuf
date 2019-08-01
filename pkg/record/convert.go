package record

import (
	"github.com/golang/protobuf/proto"
	"reflect"
)

// Convert gets proto message and converts to json like structure
func Convert(v proto.Message) map[string]interface{} {
	rv := reflect.Indirect(reflect.ValueOf(v))
	rt := rv.Type()

	converted := map[string]interface{}{}
	for i := 0; i < rt.NumField(); i++ {
		// TODO Should it pickup struct-tagged name?
		name := rt.Field(i).Name

		// Skip implicit values
		if len(name) >= 4 && name[:4] == "XXX_" {
			continue
		}

		ifVal := rv.Field(i).Interface()
		if nested, ok := ifVal.(proto.Message); ok {
			converted[name] = Convert(nested)
		} else if nestedArray, ok := ifVal.([]proto.Message); ok {
			arr := []map[string]interface{}{}
			for _, nested := range nestedArray {
				arr = append(arr, Convert(nested))
			}
			converted[name] = arr
		} else {
			converted[name] = ifVal
		}
	}

	return converted
}
