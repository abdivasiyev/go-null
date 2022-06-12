package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int64 struct {
	nullValue *sql.NullInt64
}

func (s Int64) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Int64)
	}

	return ""
}

func NewInt64(value int64) *Int64 {
	return &Int64{
		nullValue: &sql.NullInt64{
			Valid: true,
			Int64: value,
		},
	}
}

func (s Int64) Get() int64 {
	if s.nullValue != nil {
		return s.nullValue.Int64
	}

	return 0
}

func (s Int64) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Int64) Scan(value any) error {
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Int64) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullInt64{
		Int64: 0,
		Valid: false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Int64); err != nil {
			return err
		}
		s.nullValue.Valid = true
	}

	return nil
}

func (s Int64) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Int64)
	}

	return json.Marshal(nil)
}
