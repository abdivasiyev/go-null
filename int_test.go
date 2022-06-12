package null

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	testCases := []struct {
		name            string
		s               *Int
		marshalledBytes []byte
	}{
		{
			name:            "Not null marshaller",
			s:               NewInt(1),
			marshalledBytes: []byte(`1`),
		},
		{
			name:            "null marshaller",
			s:               NewInt(0),
			marshalledBytes: []byte(`0`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bytes, err := json.Marshal(tc.s)

			if err != nil {
				t.Errorf("could not marshal: %v\n", err)
				t.FailNow()
			}

			if !reflect.DeepEqual(bytes, tc.marshalledBytes) {
				t.Errorf("not equal values => got: %v, want: %v\n", string(bytes), string(tc.marshalledBytes))
				t.FailNow()
			}
		})
	}
}
