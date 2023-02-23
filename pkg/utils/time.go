package utils

import (
	"fmt"
	"time"
)

func FormatTime(value string) string {
	t, err := time.Parse("2006-01-02T15:04:05.000+08:00", value)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%d", t.Unix())
}
