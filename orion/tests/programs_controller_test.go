package programs_test

import (
    "encoding/json"
    "net/http"
	"net/http/httptest"
	"testing"
    "github.com/gin-gonic/gin"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/router"
    "github.com/ahsu1230/mathnavigatorSite/orion/test/mocks"
)

var ps mocks.ProgramService
var handler router.Handler

func init() {
	handler = router.Handler{ gin.Default(), &ps }
    handler.SetupApiEndpoints()
}

func TestHandlerGetAllPrograms(t *testing.T) {
    // Mock controller
    ps.GetAllFn = func(c *gin.Context) {
        var mockList = []*domains.Program {
            mocks.CreateMockProgram(1, "prog1", "Program1", 2, 3, "Description1"),
        }
        c.JSON(http.StatusOK, mockList)
	}

    // Send fake http to handler
    w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/programs/v1/all", nil)
	handler.Engine.ServeHTTP(w, r)

    // Validate results
    if !ps.GetAllInvoked {
		t.Fatal("expected GetAllPrograms() to be invoked")
	}
    if w.Code != http.StatusOK {
        t.Fatalf("expected code: %d, actual: %d", http.StatusOK, w.Code)
    }

    var resultList []domains.Program
    json.Unmarshal(w.Body.Bytes(), &resultList)
    if resultList[0].ProgramId != "prog1" {
        t.Fatalf("mismatched programId: %s, actual: %s", resultList[0].ProgramId, "prog1")
    }
    if resultList[0].Name != "Program1" {
        t.Fatalf("mismatched programName: %s, actual: %s", resultList[0].Name, "Program1")
    }
}

func TestHandlerGetByProgramIdSuccess(t *testing.T) {
    // Mock controller
    ps.GetByProgramIdFn = func(c *gin.Context) {
        program := mocks.CreateMockProgram(1, "prog1", "Program1", 2, 3, "Description1")
        c.JSON(http.StatusOK, program)
	}

    // Send correct fake http to handler
    w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/programs/v1/program/prog1", nil)
	handler.Engine.ServeHTTP(w, r)

    // Validate success results
    if !ps.GetByProgramIdInvoked {
		t.Fatal("expected GetAllPrograms() to be invoked")
	}
    if w.Code != http.StatusOK {
        t.Fatalf("expected code: %d, actual: %d", http.StatusOK, w.Code)
    }
    var resultProgram domains.Program
    json.Unmarshal(w.Body.Bytes(), &resultProgram)
    if resultProgram.ProgramId != "prog1" {
        t.Fatalf("mismatched programId: %s, actual: %s", resultProgram.ProgramId, "prog1")
    }
}

func TestHandlerGetByProgramIdFail(t *testing.T) {
    // Mock controller
    ps.GetByProgramIdFn = func(c *gin.Context) {
        c.Status(http.StatusNotFound)
	}

    // Send incorrect fake http to handler
    w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/programs/v1/program/asdf", nil)
	handler.Engine.ServeHTTP(w, r)

    if w.Code != http.StatusNotFound {
        t.Fatalf("expected code: %d, actual: %d", http.StatusNotFound, w.Code)
    }
}
