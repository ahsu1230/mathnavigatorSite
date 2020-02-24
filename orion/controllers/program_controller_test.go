package controllers_test

import (
    "encoding/json"
    "io"
    "net/http"
	"net/http/httptest"
	"testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/services"
    "github.com/ahsu1230/mathnavigatorSite/orion/router"
)

// Global test variables
var handler router.Handler
var mps mockProgramService

// Fake programService that implements ProgramService interface
type mockProgramService struct {
    mockGetAll func() ([]domains.Program, error)
}

// Implement methods of ProgramService interface with mocked implementations
func (mps *mockProgramService) GetAll() ([]domains.Program, error) {
	return mps.mockGetAll()
}

func init() {
	handler = router.Handler{ gin.Default() }
    handler.SetupApiEndpoints()
}

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
        t.Errorf("unexpcted error: %v\n", err)
    }
    assert.EqualValues(t, "Program1", programs[0].Name)
    assert.EqualValues(t, "prog1", programs[0].ProgramId)
    assert.EqualValues(t, "Program2", programs[1].Name)
    assert.EqualValues(t, "prog2", programs[1].ProgramId)
    assert.EqualValues(t, 2, len(programs))
}


// Helper Methods

func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
    req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
    }
    w := httptest.NewRecorder()
    handler.Engine.ServeHTTP(w, req)
    return w
}









// var ps mocks.ProgramService
// var handler router.Handler

// func init() {
// 	handler = router.Handler{ gin.Default(), &ps }
//     handler.SetupApiEndpoints()
// }

// func TestHandlerGetAllPrograms(t *testing.T) {
//     // Mock controller
//     ps.GetAllFn = func(c *gin.Context) {
//         var mockList = []*domains.Program {
//             mocks.CreateMockProgram(1, "prog1", "Program1", 2, 3, "Description1"),
//         }
//         c.JSON(http.StatusOK, mockList)
// 	}

//     // Send fake http to handler
//     w := httptest.NewRecorder()
// 	r, _ := http.NewRequest("GET", "/api/programs/v1/all", nil)
// 	handler.Engine.ServeHTTP(w, r)

//     // Validate results
//     if !ps.GetAllInvoked {
// 		t.Fatal("expected GetAllPrograms() to be invoked")
// 	}
//     if w.Code != http.StatusOK {
//         t.Fatalf("expected code: %d, actual: %d", http.StatusOK, w.Code)
//     }

//     var resultList []domains.Program
//     json.Unmarshal(w.Body.Bytes(), &resultList)
//     if resultList[0].ProgramId != "prog1" {
//         t.Fatalf("mismatched programId: %s, actual: %s", resultList[0].ProgramId, "prog1")
//     }
//     if resultList[0].Name != "Program1" {
//         t.Fatalf("mismatched programName: %s, actual: %s", resultList[0].Name, "Program1")
//     }
// }

// func TestHandlerGetByProgramIdSuccess(t *testing.T) {
//     // Mock controller
//     ps.GetByProgramIdFn = func(c *gin.Context) {
//         program := mocks.CreateMockProgram(1, "prog1", "Program1", 2, 3, "Description1")
//         c.JSON(http.StatusOK, program)
// 	}

//     // Send correct fake http to handler
//     w := httptest.NewRecorder()
// 	r, _ := http.NewRequest("GET", "/api/programs/v1/program/prog1", nil)
// 	handler.Engine.ServeHTTP(w, r)

//     // Validate success results
//     if !ps.GetByProgramIdInvoked {
// 		t.Fatal("expected GetAllPrograms() to be invoked")
// 	}
//     if w.Code != http.StatusOK {
//         t.Fatalf("expected code: %d, actual: %d", http.StatusOK, w.Code)
//     }
//     var resultProgram domains.Program
//     json.Unmarshal(w.Body.Bytes(), &resultProgram)
//     if resultProgram.ProgramId != "prog1" {
//         t.Fatalf("mismatched programId: %s, actual: %s", resultProgram.ProgramId, "prog1")
//     }
// }

// func TestHandlerGetByProgramIdFail(t *testing.T) {
//     // Mock controller
//     ps.GetByProgramIdFn = func(c *gin.Context) {
//         c.Status(http.StatusNotFound)
// 	}

//     // Send incorrect fake http to handler
//     w := httptest.NewRecorder()
// 	r, _ := http.NewRequest("GET", "/api/programs/v1/program/asdf", nil)
// 	handler.Engine.ServeHTTP(w, r)

//     if w.Code != http.StatusNotFound {
//         t.Fatalf("expected code: %d, actual: %d", http.StatusNotFound, w.Code)
//     }
// }
