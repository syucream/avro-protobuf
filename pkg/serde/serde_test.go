package serde

import (
	"testing"
	"time"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/ptypes"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		input descriptor.Message
		// TODO more validation values
	}{
		{
			input: func() descriptor.Message {
				v, _ := ptypes.TimestampProto(time.Unix(1, 2))
				return v
			}(),
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
