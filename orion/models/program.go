package models

type Program struct {
  Model                       `json:"-"`
  ProgramId       string      `gorm:"type:varchar(100)" json:"programId"`
  Name            string      `gorm:"size:255" json:"name"`
  Grade1          uint        `json:"grade1"`
  Grade2          uint        `json:"grade2"`
  Description     string      `json:"description"`
}
