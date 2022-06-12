package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int16 struct {
	nullValue *sql.NullInt16
}

func (s Int16) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Int16)
	}

	return ""
}

func NewInt16(value int16) *Int16 {
	return &Int16{
		nullValue: &sql.NullInt16{
			Valid: true,
			Int16: value,
		},
	}
}

func (s Int16) Get() int16 {
	if s.nullValue != nil {
		return s.nullValue.Int16
	}

	return 0
}

func (s Int16) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Int16) Scan(value any) error {
	s.nullValue = &sql.NullInt16{
		Int16: 0,
		Valid: false,
	}
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Int16) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullInt16{
		Int16: 0,
		Valid: false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Int16); err != nil {
			return err
		}
		s.nullValue.Valid = true
	}

	return nil
}

func (s Int16) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Int16)
	}

	return json.Marshal(nil)
}
