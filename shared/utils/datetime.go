package utils

import "time"

func GetUtcTime() time.Time {
	return time.Now().UTC()
}
