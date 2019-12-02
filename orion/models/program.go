package models

import (
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Program struct {
  gorm.Model
  ProgramId       string      `gorm:"type:varchar(100)"`
  Name            string      `gorm:"size:255"`
  Grade1          uint
  Grade2          uint
  Description     string
}
