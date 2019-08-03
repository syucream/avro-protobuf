package record

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		input    proto.Message
		expected map[string]interface{}
	}{
		{
			input: func() proto.Message {
				v, _ := ptypes.TimestampProto(time.Unix(1, 2))
				return v
			}(),
			expected: map[string]interface{}{
				"nanos":   2,
				"seconds": 1,
			},
		},
	}

	for _, c := range cases {
		actual, err := Convert(c.input)
		if err != nil {
			t.Error(err)
		}

		actualJson, actualErr := toJsonString(actual)
		if actualErr != nil {
			t.Error(actualErr)
		}

		expectedJson, expectedErr := toJsonString(c.expected)
		if expectedErr != nil {
			t.Error(expectedErr)
		}

		if actualJson != expectedJson {
			t.Errorf("expected: %v, but actual: %v", expectedJson, actualJson)
		}
	}
}

func toJsonString(schema map[string]interface{}) (string, error) {
	d, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
