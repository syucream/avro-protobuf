package schema

import (
	"encoding/json"
	"testing"

	"github.com/syucream/avro-protobuf/gen/proto/com/syucream/example"

	"github.com/linkedin/goavro"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/ptypes"
)

func TestGetRecordSchemaFromMessage(t *testing.T) {
	cases := []struct {
		input  descriptor.Message
		expect map[string]interface{}
	}{
		{
			input: ptypes.TimestampNow(),
			expect: map[string]interface{}{
				"name":      "Timestamp",
				"namespace": "google.protobuf",
				"type":      "record",
				"fields": []map[string]interface{}{
					{
						"name":    "seconds",
						"type":    "long",
						"default": 0,
					},
					{
						"name":    "nanos",
						"type":    "int",
						"default": 0,
					},
				},
			},
		},
		{
			input: &com_syucream_example.SearchResponse{},
			expect: map[string]interface{}{
				"name":      "SearchResponse",
				"namespace": "com.syucream.example",
				"type":      "record",
				"fields": []map[string]interface{}{
					{
						"default": []interface{}{},
						"name":    "results",
						"type": map[string]interface{}{
							"items": []interface{}{
								"null",
								map[string]interface{}{
									"name":      "Result",
									"namespace": "com.syucream.example",
									"type":      "record",
									"fields": []map[string]interface{}{
										{
											"default": "",
											"name":    "url",
											"type":    "string",
										},
										{
											"default": "",
											"name":    "title",
											"type":    "string",
										},
										{
											"default": []interface{}{},
											"name":    "snippets",
											"type": map[string]interface{}{
												"items": "string",
												"type":  "array",
											},
										},
									},
								},
							},
							"type": "array",
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, err := GetRecordSchemaFromMessage(c.input)
		if err != nil {
			t.Error(err)
		}

		actualJson, actualErr := toJsonString(actual)
		if actualErr != nil {
			t.Error(actualErr)
		}

		expectedJson, expectedErr := toJsonString(c.expect)
		if expectedErr != nil {
			t.Error(expectedErr)
		}

		if actualJson != expectedJson {
			t.Errorf("expected: %v, but actual: %v", expectedJson, actualJson)
		}

		if _, err = goavro.NewCodec(actualJson); err != nil {
			t.Errorf("expected Avro schema json, but it isn't: %v", err)
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
