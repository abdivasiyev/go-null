package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type String struct {
	nullValue *sql.NullString
}

func (s String) String() string {
	if s.nullValue != nil {
		return s.nullValue.String
	}

	return ""
}

func NewString(value string) String {
	return String{
		nullValue: &sql.NullString{
			Valid:  value != "",
			String: value,
		},
	}
}

func (s String) Get() string {
	if s.nullValue != nil {
		return s.nullValue.String
	}

	return ""
}

func (s *String) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *String) Scan(value any) error {
	if err := s.nullValue.Scan(value); err != nil {
		return err
	}

	return nil
}

func (s *String) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullString{
		String: "",
		Valid:  false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.String); err != nil {
			return err
		}
	}

	return nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid && s.nullValue.String != "" {
		return json.Marshal(s.nullValue.String)
	}

	return json.Marshal(nil)
}
