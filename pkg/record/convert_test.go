package record

import (
	"github.com/syucream/avro-protobuf/gen/proto/com/syucream/example"
	"reflect"
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
		{
			input: &com_syucream_example.SearchResponse{
				Results: []*com_syucream_example.SearchResponse_Result{
					{
						Url:   "http://example.com",
						Title: "title",
						Snippets: []string{
							"snippet",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"results": []map[string]interface{}{
					{
						"url":   "http://example.com",
						"title": "title",
						"snippets": []string{
							"snippet",
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, err := Convert(c.input)
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(actual, c.expected) {
			t.Errorf("expected: %v, but actual: %v", c.expected, actual)
		}
	}
}
