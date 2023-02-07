package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().String())
	d, err := time.ParseDuration("1675749670")
	//tm, err := time.Parse("2006-01-02 15:04:05", "1675749670")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d.String())
	fmt.Println(time.Unix(1675749670, 0))
	fmt.Println(time.Unix(1675749670, 0).Local())
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-02-07 15:11:50", time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1.Unix())

	fmt.Println(time.Unix(1675754584259, 1e6))
	fmt.Println(time.UnixMilli(1675754584259))
	fmt.Println(time.Unix(time.UnixMilli(1675754584259).Unix(), 0))
	fmt.Println(time.Unix(time.UnixMilli(1675737704).Unix(), 0))
}
