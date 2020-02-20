package domains

import (
	"database/sql"
    "github.com/gin-gonic/gin"
)

type Program struct {
	Id          uint
	CreatedAt   uint          `db:"created_at"`
	UpdatedAt   uint          `db:"updated_at"`
	DeletedAt   sql.NullInt64 `db:"deleted_at"`
	ProgramId   string        `db:"program_id" json:"programId"`
	Name        string        `json:"name"`
	Grade1      uint          `json:"grade1"`
	Grade2      uint          `json:"grade2"`
	Description string        `json:"description"`
}

type ProgramService interface {
    GetAll(c *gin.Context)
    GetByProgramId(c *gin.Context)
    Create(c *gin.Context)
    Update(c *gin.Context)
    Delete(c *gin.Context)
}
