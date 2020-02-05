package utils

import (
  "time"
)

func TimestampNow() int64 {
    return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}
