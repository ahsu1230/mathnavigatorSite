package domains

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Alias data types for handling sql.Nullxxx & JSON Marshaling

// NullString - alias data type wrapper around sql.NullString
type NullString struct {
	sql.NullString
}

func NewNullString(str string) NullString {
	return NullString{sql.NullString{String: str, Valid: str != ""}}
}

// MarshalJSON for NullString
func (nullString *NullString) MarshalJSON() ([]byte, error) {
	if !nullString.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nullString.String)
}

// UnmarshalJSON for NullString
func (nullString *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nullString.String)
	nullString.Valid = err == nil && nullString.String != ""
	return err
}

// NullUint represents an uint that may be null.
// NullUint implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullUint struct {
	Uint  uint
	Valid bool // Valid is true if Uint is not NULL
}

func NewNullUint(num uint) NullUint {
	return NullUint{
		Uint:  num,
		Valid: num != 0,
	}
}

// Scan implements the Scanner interface.
func (n *NullUint) Scan(value interface{}) error {
	num, ok := value.(int64)
	valid := ok && num != 0
	*n = NullUint{uint(num), valid}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint), nil
}

// MarshalJSON for NullUint
func (n *NullUint) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Uint)
}

// UnmarshalJSON for NullUint
func (n *NullUint) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &n.Uint)
	n.Valid = err == nil && n.Uint != 0
	return err
}
