package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	nullValue *sql.NullTime
}

func (s Time) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Time)
	}

	return ""
}

func NewTime(value time.Time) *Time {
	return &Time{
		nullValue: &sql.NullTime{
			Valid: true,
			Time:  value,
		},
	}
}

func (s Time) Get() time.Time {
	if s.nullValue != nil {
		return s.nullValue.Time
	}

	return time.Time{}
}

func (s Time) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Time) Scan(value any) error {
	s.nullValue = &sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Time) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Time); err != nil {
			return err
		}
		s.nullValue.Valid = true
	}

	return nil
}

func (s Time) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Time)
	}

	return json.Marshal(nil)
}
