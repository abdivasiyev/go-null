package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Bool struct {
	nullValue *sql.NullBool
}

func (s Bool) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Bool)
	}

	return ""
}

func NewBool(value bool) Bool {
	return Bool{
		nullValue: &sql.NullBool{
			Valid: true,
			Bool:  value,
		},
	}
}

func (s Bool) Get() bool {
	if s.nullValue != nil {
		return s.nullValue.Bool
	}

	return false
}

func (s *Bool) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Bool) Scan(value any) error {
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *Bool) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullBool{
		Bool:  false,
		Valid: false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Bool); err != nil {
			return err
		}
	}

	return nil
}

func (s Bool) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Bool)
	}

	return json.Marshal(nil)
}
