package serde

import (
	"github.com/golang/protobuf/descriptor"
	"github.com/syucream/avro-protobuf/gen/proto/com/syucream/example"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		input descriptor.Message
		// TODO more validation values
	}{
		{
			input: &com_syucream_example.SearchResponse{
				Results: []*com_syucream_example.SearchResponse_Result{
					{
						Url: "http://example.com",
						Title: "title",
						Snippets: []string{
							"snippet",
						},
					},
				},
			},
		},
		// TODO more cases
	}

	for _, c := range cases {
		serDe, err := NewSerDe(c.input)
		if err != nil {
			t.Error(err)
		}

		bin, err := serDe.Serialize(c.input)
		if err != nil {
			t.Error(err)
		}

		v := c.input
		err = serDe.Deserialize(bin, v)
		if err != nil {
			t.Error(err)
		}
	}
}
