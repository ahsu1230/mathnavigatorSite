package sql_helper

import (
	"database/sql/driver"
)

// NullUint represents an uint that may be null.
type NullUint struct {
	Uint  uint
	Valid bool // Valid is true if Uint is not NULL
}

// Value implements the driver Valuer interface.
func (n NullUint) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint, nil
}
