package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Float32 struct {
	nullValue *sql.NullFloat64
}

func (s Float32) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Float64)
	}

	return ""
}

func NewFloat32(value float32) *Float32 {
	return &Float32{
		nullValue: &sql.NullFloat64{
			Valid:   true,
			Float64: float64(value),
		},
	}
}

func (s Float32) Get() float32 {
	if s.nullValue != nil {
		return float32(s.nullValue.Float64)
	}

	return 0
}

func (s Float32) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Float32) Scan(value any) error {
	s.nullValue = &sql.NullFloat64{
		Valid:   false,
		Float64: 0,
	}
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Float32) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullFloat64{
		Float64: 0,
		Valid:   false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Float64); err != nil {
			return err
		}
		s.nullValue.Valid = true
	}

	return nil
}

func (s Float32) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Float64)
	}

	return json.Marshal(nil)
}
