package testUtils

import (
	"context"
	"database/sql/driver"
	"time"
)

var TimeNow = time.Now().UTC()
var TimeLater = TimeNow.Add(time.Hour * 24 * 30) // One month later
var Context context.Context = context.TODO()

// For sqlmock - a matcher for any time
type MockAnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a MockAnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
