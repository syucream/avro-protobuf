package serde

import (
	"fmt"

	"github.com/syucream/avro-protobuf/pkg/record"
)

var (
	ErrBrokenGoavroValue  = fmt.Errorf("goavro native value is broken")
	ErrInvalidGoavroUnion = fmt.Errorf("It is not a goavro union value")
)

func unwrapUnion(orig map[string]interface{}) (map[string]interface{}, error) {
	// An actual union value in goavro must have only 1 element
	if len(orig) == 1 {
		unwrapped := map[string]interface{}{}

		for k, v := range orig {
			// And the key is union key
			if record.IsUnionKey(k) {
				if vv, ok := v.(map[string]interface{}); ok {
					unwrapped = vv
				} else {
					return nil, ErrInvalidGoavroUnion
				}
			}
		}

		return unwrapped, nil
	}

	return orig, nil
}

func toProtoJsonArray(orig []interface{}) ([]interface{}, error) {
	arr := []interface{}{}

	for _, v := range orig {
		vv, err := toProtoJson(v)
		if err != nil {
			return nil, err
		}
		arr = append(arr, vv)
	}

	return arr, nil
}

// TODO rename?
func toProtoJson(orig interface{}) (map[string]interface{}, error) {
	switch orig.(type) {
	case map[string]interface{}:
		m, err := unwrapUnion(orig.(map[string]interface{}))
		if err != nil {
			return nil, err
		}

		protoJson := map[string]interface{}{}
		for k, v := range m {
			switch v.(type) {

			case int64, uint64:
				protoJson[k] = fmt.Sprint(v)

			case map[string]interface{}:
				sub, err := toProtoJson(v.(map[string]interface{}))
				if err != nil {
					return nil, err
				}
				protoJson[k] = sub

			case []interface{}:
				sub, err := toProtoJsonArray(v.([]interface{}))
				if err != nil {
					return nil, err
				}
				protoJson[k] = sub

			// TODO timestamp

			// TODO duration

			default:
				protoJson[k] = v
			}
		}

		return protoJson, nil

	// TODO array??

	default:
		return nil, ErrBrokenGoavroValue
	}
}
