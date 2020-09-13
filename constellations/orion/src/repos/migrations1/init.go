package migrations1

import (
	"fmt"
	"github.com/pressly/goose"
)

func init() {
	fmt.Println("INVOKE MIGRATION!!!!!!")
	goose.AddMigration(Up00001, Down00001)
}