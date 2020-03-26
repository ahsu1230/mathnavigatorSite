package domains_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"testing"
)

func TestNullString(t *testing.T) {
	ns1 := domains.CreateNullString("asdf")
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
	ns1 := domains.CreateNullString("")
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