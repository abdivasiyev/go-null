package null

import (
	"encoding/json"
	"testing"
)

func TestFloat32(t *testing.T) {
	testCases := []struct {
		name string
		s    *Float32
	}{
		{
			name: "Not null marshaller",
			s:    NewFloat32(3.2),
		},
		{
			name: "null marshaller",
			s:    NewFloat32(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := json.Marshal(tc.s)

			if err != nil {
				t.Errorf("could not marshal: %v\n", err)
				t.FailNow()
			}
		})
	}
}
