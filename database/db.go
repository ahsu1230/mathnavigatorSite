package database

import (
  "fmt"

  "github.com/jmoiron/sqlx"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/lib/pq"
)

var DbConn *sqlx.DB

func createConnectionInfo(host string, port int, user string, pass string) (string) {
  dbSchema := "mathnavdb"
  info := fmt.Sprintf("%s:%s@(%s)/%s", user, pass, host, dbSchema)
  info += "?charset=utf8&parseTime=True&loc=Local"
  fmt.Println(info)
  return info
}

func OpenDb(host string, port int, user string, pass string) {
  var err error
  connection := createConnectionInfo(host, port, user, pass)
  DbConn, err = sqlx.Connect("mysql", connection)
  if err != nil {
    panic(err)
  }
}

func CloseDb() {
  DbConn.Close()
}
