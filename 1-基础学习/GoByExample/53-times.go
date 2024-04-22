package main

import (
	"fmt"
	"time"
)

func main() {
	// 简写
	p := fmt.Println

	// 获取当前时间
	now := time.Now()
	p(now)

	// 创建一个时间，格式为 2009-11-17 20:34:58.651387237 +0000 UTC m=+0.000000001
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	// then.Before(now) 判断 then 是否在 now 之前
	p(then.Before(now))
	// then.After(now) 判断 then 是否在 now 之后
	p(then.After(now))
	// then.Equal(now) 判断 then 是否等于 now
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add 将时间往后移动一个时间段
	p(then.Add(diff))
	p(then.Add(-diff))

}
