package null

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	testCases := []struct {
		name            string
		s               String
		marshalledBytes []byte
	}{
		{
			name:            "Not null marshaller",
			s:               NewString("test"),
			marshalledBytes: []byte(`"test"`),
		},
		{
			name:            "null marshaller",
			s:               NewString(""),
			marshalledBytes: []byte(`null`),
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
