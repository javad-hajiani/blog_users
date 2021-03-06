package dateTime

import "time"

const apiLayout = "2006-01-02T15:04:05"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiLayout)
}

