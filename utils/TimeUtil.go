package utils

import (
	"strconv"
	"time"
)

func ParseStringTime(timeStr string) (time.Time, error) {
	timeUnix, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	//unix := time.Unix(time.Now().Unix(), 0)
	//fmt.Println(unix)
	return time.UnixMilli(timeUnix), nil
}
