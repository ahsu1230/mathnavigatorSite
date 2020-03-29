package domains_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
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
	fmt.Println("json", string(bytes))

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
	fmt.Println("json", string(bytes))

	var result domains.NullString
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, "", result.String)
	assert.EqualValues(t, false, result.Valid)
}

func TestNullUint(t *testing.T) {
	nu1 := domains.NewNullUint(14)
	bytes, err := json.Marshal(&nu1)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	fmt.Println("json", string(bytes))

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
	fmt.Println("json", string(bytes))

	var result domains.NullUint
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	assert.EqualValues(t, uint(0), result.Uint)
	assert.EqualValues(t, false, result.Valid)
}
