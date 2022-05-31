package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Float64 struct {
	nullValue *sql.NullFloat64
}

func (s Float64) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Float64)
	}

	return ""
}

func NewFloat64(value float64) Float64 {
	return Float64{
		nullValue: &sql.NullFloat64{
			Valid:   true,
			Float64: value,
		},
	}
}

func (s Float64) Get() float64 {
	if s.nullValue != nil {
		return s.nullValue.Float64
	}

	return 0
}

func (s *Float64) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Float64) Scan(value any) error {
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Float64) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullFloat64{
		Float64: 0,
		Valid:   false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Float64); err != nil {
			return err
		}
	}

	return nil
}

func (s Float64) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Float64)
	}

	return json.Marshal(nil)
}
