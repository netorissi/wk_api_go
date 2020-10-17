package utils

import "time"

func DateNow() int64 {
	return time.Now().UnixNano() / 1e6
}
