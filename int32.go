package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int32 struct {
	nullValue *sql.NullInt32
}

func (s Int32) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Int32)
	}

	return ""
}

func NewInt32(value int32) *Int32 {
	return &Int32{
		nullValue: &sql.NullInt32{
			Valid: true,
			Int32: value,
		},
	}
}

func (s Int32) Get() int32 {
	if s.nullValue != nil {
		return s.nullValue.Int32
	}

	return 0
}

func (s Int32) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Int32) Scan(value any) error {
	s.nullValue = &sql.NullInt32{
		Int32: 0,
		Valid: false,
	}
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Int32) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Int32); err != nil {
			return err
		}
		s.nullValue.Valid = true
	}

	return nil
}

func (s Int32) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Int32)
	}

	return json.Marshal(nil)
}
