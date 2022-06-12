package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Byte struct {
	nullValue *sql.NullByte
}

func (s Byte) String() string {
	if s.nullValue != nil {
		return fmt.Sprint(s.nullValue.Byte)
	}

	return ""
}

func NewByte(value byte) *Byte {
	return &Byte{
		nullValue: &sql.NullByte{
			Valid: true,
			Byte:  value,
		},
	}
}

func (s Byte) Get() byte {
	if s.nullValue != nil {
		return s.nullValue.Byte
	}

	return 0
}

func (s Byte) Value() (driver.Value, error) {
	if s.nullValue != nil {
		return s.nullValue.Value()
	}

	return nil, nil
}

func (s *Byte) Scan(value any) error {
	s.nullValue = &sql.NullByte{
		Byte:  0,
		Valid: false,
	}

	return s.nullValue.Scan(value)
}

func (s *Byte) UnmarshalJSON(bytes []byte) error {
	s.nullValue = &sql.NullByte{
		Byte:  0,
		Valid: false,
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &s.nullValue.Byte); err != nil {
			return err
		}
		s.nullValue.Valid = true
	}

	return nil
}

func (s Byte) MarshalJSON() ([]byte, error) {
	if s.nullValue.Valid {
		return json.Marshal(s.nullValue.Byte)
	}

	return json.Marshal(nil)
}
