package programs

import (
  "database/sql"
)

type Program struct {
  Id              uint
  CreatedAt       uint        `db:"created_at"`
  UpdatedAt       uint        `db:"updated_at"`
  DeletedAt       sql.NullInt64 `db:"deleted_at"`
  ProgramId       string      `db:"program_id" json:"programId"`
  Name            string      `json:"name"`
  Grade1          uint        `json:"grade1"`
  Grade2          uint        `json:"grade2"`
  Description     string      `json:"description"`
}
