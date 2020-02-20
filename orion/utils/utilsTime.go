package utils

import (
	"time"
)

// Create timestamp in milliseconds
func TimestampNow() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
