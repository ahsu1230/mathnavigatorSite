package controllers_test

import (
    "bytes"
    "encoding/json"
    "errors"
    "io"
    "net/http"
	"net/http/httptest"
	"testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
    "github.com/ahsu1230/mathnavigatorSite/orion/pkg/router"
)

// Global test variables
var handler router.Handler
var mps mockProgramService

// Fake programService that implements ProgramService interface
type mockProgramService struct {
    mockGetAll func() ([]domains.Program, error)
    mockGetByProgramId func(string) (domains.Program, error)
    mockCreate func(domains.Program) error
    mockUpdate func(string, domains.Program) error
    mockDelete func(string) error
}

// Implement methods of ProgramService interface with mocked implementations
func (mps *mockProgramService) GetAll() ([]domains.Program, error) {
	return mps.mockGetAll()
}
func (mps *mockProgramService) GetByProgramId(programId string) (domains.Program, error) {
	return mps.mockGetByProgramId(programId)
}
func (mps *mockProgramService) Create(program domains.Program) error {
	return mps.mockCreate(program)
}
func (mps *mockProgramService) Update(programId string, program domains.Program) error {
	return mps.mockUpdate(programId, program)
}
func (mps *mockProgramService) Delete(programId string) error {
	return mps.mockDelete(programId)
}

func init() {
    gin.SetMode(gin.TestMode)
    engine := gin.Default()
	handler = router.Handler{ Engine: engine }
    handler.SetupApiEndpoints()
}

//
// Test Get All
//
func TestGetAllPrograms_Success(t *testing.T) {
	mps.mockGetAll = func() ([]domains.Program, error) {
		 return []domains.Program{
			{
				Id:         1,
				ProgramId:  "prog1",
                Name:       "Program1",
                Grade1:     2,
                Grade2:     3,
                Description:    "Description1",
			},
			{
				Id:         2,
				ProgramId:  "prog2",
                Name:       "Program2",
                Grade1:     2,
                Grade2:     3,
                Description:    "Description2",
			},
		}, nil
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    recorder := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/all", nil)

    // Validate results
    assert.EqualValues(t, http.StatusOK, recorder.Code)
    var programs []domains.Program
    if err := json.Unmarshal(recorder.Body.Bytes(), &programs); err != nil {
        t.Errorf("unexpected error: %v\n", err)
    }
    assert.EqualValues(t, "Program1", programs[0].Name)
    assert.EqualValues(t, "prog1", programs[0].ProgramId)
    assert.EqualValues(t, "Program2", programs[1].Name)
    assert.EqualValues(t, "prog2", programs[1].ProgramId)
    assert.EqualValues(t, 2, len(programs))
}

//
// Test Get Program
//
func TestGetProgram_Success(t *testing.T) {
	mps.mockGetByProgramId = func(programId string) (domains.Program, error) {
        program := createMockProgram("prog1", "Program1", 2, 3, "descript1")
        return program, nil
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    recorder := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog1", nil)

    // Validate results
    assert.EqualValues(t, http.StatusOK, recorder.Code)
    var program domains.Program
    if err := json.Unmarshal(recorder.Body.Bytes(), &program); err != nil {
        t.Errorf("unexpected error: %v\n", err)
    }
    assert.EqualValues(t, "prog1", program.ProgramId)
    assert.EqualValues(t, "Program1", program.Name)
}

func TestGetProgram_Failure(t *testing.T) {
	mps.mockGetByProgramId = func(programId string) (domains.Program, error) {
		 return domains.Program{}, errors.New("Not Found")
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    recorder := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog2", nil)

    // Validate results
    assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateProgram_Success(t *testing.T) {
	mps.mockCreate = func(program domains.Program) error {
		 return nil
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    program := createMockProgram("prog1", "Program1", 2, 3, "descript1")
    marshal, _ := json.Marshal(program)
    body := bytes.NewBuffer(marshal)
    recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body)

    // Validate results
    assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateProgram_Failure(t *testing.T) {
    // no mock needed
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    program := createMockProgram("prog1", "", 2, 3, "descript1") // Empty Name!
    marshal, _ := json.Marshal(program)
    body := bytes.NewBuffer(marshal)
    recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body)

    // Validate results
    assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateProgram_Success(t *testing.T) {
	mps.mockUpdate = func(programId string, program domains.Program) error {
		 return nil // Succesful update
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    program := createMockProgram("prog2", "Program2", 2, 3, "descript2")
    body := createBodyFromProgram(program)
    recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", body)

    // Validate results
    assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateProgram_Invalid(t *testing.T) {
    // no mock needed
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    program := createMockProgram("prog2", "", 2, 3, "descript2") // Empty Name!
    body := createBodyFromProgram(program)
    recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", body)

    // Validate results
    assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateProgram_Failure(t *testing.T) {
	mps.mockUpdate = func(programId string, program domains.Program) error {
		 return errors.New("not found")
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    program := createMockProgram("prog2", "Program2", 2, 3, "descript2")
    body := createBodyFromProgram(program)
    recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", body)

    // Validate results
    assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteProgram_Success(t *testing.T) {
	mps.mockDelete = func(programId string) error {
		return nil // Return no error, successful delete!
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    recorder := sendHttpRequest(t, http.MethodDelete, "/api/programs/v1/program/some_program", nil)

    // Validate results
    assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteProgram_Failure(t *testing.T) {
	mps.mockDelete = func(programId string) error {
		return errors.New("not found")
    }
    services.ProgramService = &mps
    
    // Create new HTTP request to endpoint
    recorder := sendHttpRequest(t, http.MethodDelete, "/api/programs/v1/program/some_program", nil)

    // Validate results
    assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockProgram(programId string, name string, grade1 uint, grade2 uint, description string) domains.Program {
    return domains.Program{
        ProgramId: programId,
        Name: name,
        Grade1: grade1,
        Grade2: grade2,
        Description: description,
    }
}

func createBodyFromProgram(program domains.Program) io.Reader {
    marshal, err := json.Marshal(program)
    if err != nil {
        panic(err)
    }
    return bytes.NewBuffer(marshal)
}

func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
    req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
    }
    w := httptest.NewRecorder()
    handler.Engine.ServeHTTP(w, req)
    return w
}