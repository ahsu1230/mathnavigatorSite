package models
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var openDb *gorm.DB

func OpenDb() {
  db, err := gorm.Open("sqlite3", "tmp/gorm.db")
  if err != nil {
    panic("failed to connect database")
  }

  openDb = db
  openDb.AutoMigrate(&Program{})
}

func GetDb() (*gorm.DB) {
  return openDb
}

func CloseDb() {
  openDb.Close()
}
