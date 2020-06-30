package integration_tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Users and GetAll()
func Test_CreateUsers(t *testing.T) {
	family1 := createFamily(1)
	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	family2 := createFamily(2)
	body2 := createJsonBody(&family2)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	createAllUsers(t)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/family/1", nil)

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/family/2", nil)
	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, 1, users[0])
	assert.EqualValues(t, 1, len(users))

	if err := json.Unmarshal(recorder3.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 2, users[0])
	assertUser(t, 3, users[1])
	assert.EqualValues(t, 2, len(users))

	resetTable(t, domains.TABLE_USERS)
	resetTable(t, domains.TABLE_FAMILIES)

}

// Test: Create 3 Users and search by pagination
// func Test_SearchUsers(t *testing.T) {
// 	createAllUsers(t)

// 	// Call Get All Searching for "Smith" With Page Size 2 Offset 0
// 	recorder1 := sendHttpRequest(t, http.MethodGet, "/api/users/all?search=Smith&pageSize=2&offset=0", nil)

// 	// Validate results
// 	assert.EqualValues(t, http.StatusOK, recorder1.Code)
// 	var users1 []domains.User
// 	if err := json.Unmarshal(recorder1.Body.Bytes(), &users1); err != nil {
// 		t.Errorf("unexpected error: %v\n", err)
// 	}
// 	assertUser(t, 1, users1[0])
// 	assertUser(t, 2, users1[1])
// 	assert.EqualValues(t, 2, len(users1))

// 	// Call Get All Searching for "Smith" With Page Size 2 Offset 2
// 	recorder2 := sendHttpRequest(t, http.MethodGet, "/api/users/all?search=Smith&pageSize=2&offset=2", nil)

// 	// Validate results
// 	assert.EqualValues(t, http.StatusOK, recorder2.Code)
// 	var users2 []domains.User
// 	if err := json.Unmarshal(recorder2.Body.Bytes(), &users2); err != nil {
// 		t.Errorf("unexpected error: %v\n", err)
// 	}
// 	assertUser(t, 3, users2[0])
// 	assert.EqualValues(t, 1, len(users2))

// 	resetTable(t, domains.TABLE_USERS)
// }

// Test: Create 3 Users and GetUserByFamilyId
func Test_GetUsersByFamilyId(t *testing.T) {
	family1 := createFamily(1)
	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	family2 := createFamily(2)
	body2 := createJsonBody(&family2)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	createAllUsers(t)
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/family/1", nil)

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/family/2", nil)
	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, 1, users[0])
	assert.EqualValues(t, 1, len(users))

	if err := json.Unmarshal(recorder3.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 2, users[0])
	assertUser(t, 3, users[1])
	assert.EqualValues(t, 2, len(users))

	resetTable(t, domains.TABLE_USERS)
	resetTable(t, domains.TABLE_FAMILIES)

}

// Test: Create 1 Family, 1 User, Update it, GetUserById()
func Test_UpdateUser(t *testing.T) {
	family := createFamily(1)
	body := createJsonBody(&family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Create 1 User
	user1 := createUser(1)
	body1 := createJsonBody(&user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedUser := createUser(4)
	updatedBody := createJsonBody(&updatedUser)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/users/user/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var user domains.User
	if err := json.Unmarshal(recorder3.Body.Bytes(), &user); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, 4, user)

	resetTable(t, domains.TABLE_USERS)
	resetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 1 User, Delete it, GetByUserId()
func Test_DeleteUser(t *testing.T) {
	family := createFamily(1)
	body := createJsonBody(&family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Create
	user1 := createUser(1)
	body1 := createJsonBody(&user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetTable(t, domains.TABLE_USERS)
	resetTable(t, domains.TABLE_FAMILIES)

}

// Helper methods
func createUser(id int) domains.User {
	switch id {
	case 1:
		return domains.User{
			FirstName:  "John",
			LastName:   "Smith",
			MiddleName: domains.NewNullString("Middle"),
			Email:      "john_smith@example.com",
			Phone:      "555-555-0100",
			IsGuardian: true,
			FamilyId:   1,
			Notes:      domains.NewNullString("notes1"),
		}
	case 2:
		return domains.User{
			FirstName:  "Bob",
			LastName:   "Smith",
			MiddleName: domains.NewNullString(""),
			Email:      "bob_smith@example.com",
			Phone:      "555-555-0101",
			IsGuardian: false,
			FamilyId:   2,
			Notes:      domains.NewNullString("notes2"),
		}
	case 3:
		return domains.User{
			FirstName:  "Foo",
			LastName:   "Bar",
			MiddleName: domains.NewNullString("Smith"),
			Email:      "foobar@example.com",
			Phone:      "555-555-0102",
			IsGuardian: false,
			FamilyId:   2,
			Notes:      domains.NewNullString("notes3"),
		}
	case 4:
		return domains.User{
			FirstName:  "Austin",
			LastName:   "Hsu",
			MiddleName: domains.NewNullString(""),
			Email:      "austinhsu@example.com",
			Phone:      "555-555-0103",
			IsGuardian: false,
			FamilyId:   1,
			Notes:      domains.NewNullString("notes4"),
		}
	default:
		return domains.User{}
	}
}

func createAllUsers(t *testing.T) {
	for i := 1; i < 4; i++ {
		user := createUser(i)
		body := createJsonBody(&user)
		recorder := sendHttpRequest(t, http.MethodPost, "/api/users/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func assertUser(t *testing.T, id int, user domains.User) {
	switch id {
	case 1:
		assert.EqualValues(t, "John", user.FirstName)
		assert.EqualValues(t, "Smith", user.LastName)
		assert.EqualValues(t, "Middle", user.MiddleName.String)
		assert.EqualValues(t, "john_smith@example.com", user.Email)
		assert.EqualValues(t, "555-555-0100", user.Phone)
		assert.EqualValues(t, true, user.IsGuardian)
		assert.EqualValues(t, 1, user.FamilyId)
		assert.EqualValues(t, "notes1", user.Notes.String)
	case 2:
		assert.EqualValues(t, "Bob", user.FirstName)
		assert.EqualValues(t, "Smith", user.LastName)
		assert.EqualValues(t, "Middle", user.MiddleName.String)
		assert.EqualValues(t, "bob_smith@example.com", user.Email)
		assert.EqualValues(t, "555-555-0101", user.Phone)
		assert.EqualValues(t, false, user.IsGuardian)
		assert.EqualValues(t, 2, user.FamilyId)
		assert.EqualValues(t, "notes2", user.Notes.String)

	case 3:
		assert.EqualValues(t, "Foo", user.FirstName)
		assert.EqualValues(t, "Bar", user.LastName)
		assert.EqualValues(t, "Smith", user.MiddleName.String)
		assert.EqualValues(t, "foobar@example.com", user.Email)
		assert.EqualValues(t, "555-555-0102", user.Phone)
		assert.EqualValues(t, false, user.IsGuardian)
		assert.EqualValues(t, 2, user.FamilyId)
		assert.EqualValues(t, "notes3", user.Notes.String)

	}

}

// func createBodyFromFamily(family domains.Family) io.Reader {
// 	marshal, err := json.Marshal(&family)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return bytes.NewBuffer(marshal)
// }
