package domains_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/stretchr/testify/assert"
)

func TestNullString(t *testing.T) {
	ns1 := domains.NewNullString("asdf")
	assert.EqualValues(t, "asdf", ns1.String)
	assert.EqualValues(t, true, ns1.Valid)

	bytes, err := json.Marshal(&ns1)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	t.Log("json", string(bytes))

	var result domains.NullString
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, "asdf", result.String)
	assert.EqualValues(t, true, result.Valid)
}

func TestNullStringEmpty(t *testing.T) {
	ns1 := domains.NewNullString("")
	bytes, err := json.Marshal(&ns1)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	t.Log("json", string(bytes))

	var result domains.NullString
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, "", result.String)
	assert.EqualValues(t, false, result.Valid)
}

func TestNullTime(t *testing.T) {
	now := time.Now().UTC()
	nullTime := domains.NewNullTime(now)
	assert.EqualValues(t, now, nullTime.Time)
	assert.EqualValues(t, true, nullTime.Valid)

	bytes, err := json.Marshal(&nullTime)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	t.Log("json", string(bytes))

	var result domains.NullTime
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, now, result.Time)
	assert.EqualValues(t, true, result.Valid)
}

func TestNullTimeEmpty(t *testing.T) {
	zeroTime := time.Time{}
	nullTime := domains.NewNullTime(zeroTime)
	bytes, err := json.Marshal(&nullTime)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	t.Log("json", string(bytes))

	var result domains.NullTime
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, zeroTime, result.Time)
	assert.EqualValues(t, false, result.Valid)
}

func TestNullUint(t *testing.T) {
	nu1 := domains.NewNullUint(14)
	bytes, err := json.Marshal(&nu1)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	t.Log("json", string(bytes))

	var result domains.NullUint
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, uint(14), result.Uint)
	assert.EqualValues(t, true, result.Valid)
}

func TestNullUintEmpty(t *testing.T) {
	nu1 := domains.NewNullUint(0)
	bytes, err := json.Marshal(&nu1)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	t.Log("json", string(bytes))

	var result domains.NullUint
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, uint(0), result.Uint)
	assert.EqualValues(t, false, result.Valid)
}
