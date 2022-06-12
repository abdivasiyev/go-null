package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int struct {
	nullValue *sql.NullInt32
}

func (s Int) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Int32)
	}

	return ""
}

func NewInt(value int) *Int {
	return &Int{
		nullValue: &sql.NullInt32{
			Valid: true,
			Int32: int32(value),
		},
	}
}

func (s Int) Get() int {
	if s.nullValue != nil {
		return int(s.nullValue.Int32)
	}

	return 0
}

func (s Int) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Int) Scan(value any) error {
	s.nullValue = &sql.NullInt32{
		Int32: 0,
		Valid: false,
	}
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Int) UnmarshalJSON(bytes []byte) error {
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

func (s Int) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Int32)
	}

	return json.Marshal(nil)
}
