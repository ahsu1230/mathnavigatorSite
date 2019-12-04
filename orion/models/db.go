package models
import (
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var openDb *gorm.DB

func OpenDb() {
  db, err := gorm.Open("mysql", "root@/db?charset=utf8&parseTime=True&loc=Local")
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

// Custom Model instead of using gorm.Model
type Model struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *time.Time `sql:"index"`
}

func (m *Model) BeforeCreate(scope *gorm.Scope) error {
  now := time.Now().Unix()
	scope.SetColumn("CreatedAt", now)
  if m.UpdatedAt == 0 {
    scope.SetColumn("UpdatedAt", now)
  }
	return nil
}

func (m *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
