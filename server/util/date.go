package util

import "time"

func Now() string {
	const LAYOUT = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(LAYOUT)
}
