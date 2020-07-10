package testUtils

import (
	"time"
)

var TimeNow = time.Now().UTC()
var TimeLater = TimeNow.Add(time.Hour * 24 * 30) // One month later
