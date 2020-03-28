package sql_helper

import (
	"database/sql/driver"
)

// NullUint represents an uint that may be null.
// NullUint implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullUint struct {
	Uint  uint
	Valid bool // Valid is true if Uint is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint) Scan(value interface{}) error {
	n.Uint, n.Valid = value.(uint)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint, nil
}
