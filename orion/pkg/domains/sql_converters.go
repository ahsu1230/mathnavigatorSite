package domains

import (
	"database/sql"
	"encoding/json"
)

// Alias data types for handling sql.Nullxxx & JSON Marshaling 

type NullString struct {
	sql.NullString
}

func CreateNullString(str string) NullString {
	return NullString{ sql.NullString{String: str, Valid: str != ""} }
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