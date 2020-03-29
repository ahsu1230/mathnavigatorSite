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

func CreateNullString(str string) NullString {
	return NullString{sql.NullString{String: str, Valid: str != ""}}
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil && ns.String != "")
	return err
}

// NullUint represents an uint that may be null.
// NullUint implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullUint struct {
	Num   uint
	Valid bool // Valid is true if Num is not NULL
}

func CreateNullUint(num uint) NullUint {
	return NullUint{
		Num:   num,
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
	return int64(n.Num), nil
}

// MarshalJSON for NullUint
func (n *NullUint) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Num)
}

// UnmarshalJSON for NullUint
func (n *NullUint) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &n.Num)
	n.Valid = (err == nil && n.Num != 0)
	return err
}
