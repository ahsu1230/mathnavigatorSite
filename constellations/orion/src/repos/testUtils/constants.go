package testUtils

import (
	"context"
	"time"
)

var TimeNow = time.Now().UTC()
var TimeLater = TimeNow.Add(time.Hour * 24 * 30) // One month later
var Context context.Context = context.TODO()
