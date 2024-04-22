package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	// Unix 返回 t 表示的时间，使用 Unix 时间，即从时间点 January 1, 1970 UTC 到时间点 t 所经过的时间（单位秒）。
	secs := now.Unix()
	// UnixNano 返回 t 表示的时间，使用 Unix 时间，即从时间点 January 1, 1970 UTC 到时间点 t 所经过的时间（单位纳秒）。
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
