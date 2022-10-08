package gopkg

import (
	"fmt"
	"time"
)

// now := time.Now() 	  // 获取当前时间
// year := now.Year()     // 年
// month := now.Month()   // 月
// day := now.Day()       // 日
// hour := now.Hour()     // 时
// minute := now.Minute() // 分
// second := now.Second() // 秒
// week := now.Weekday()  // 周
// _, weekTh := now.ISOWeek()	// 第几周
// timestamp := now.Unix() 		// 时间戳
// time.Unix(timestamp, 0) 		// 将时间戳转为时间格式
// 时间加减：Add, Sub
// 时间比较：Equal, Before, After
// later := now.Add(time.Hour)  // 当前时间加1小时后的时间

// TimeToString
// 2022-10-07 15:05:12
func TimeToString() {
	t := time.Now()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}

// TimeFormat
// 2022/10/07 15:05:12 PM Fri Oct
func TimeFormat() string {
	return time.Now().Format("2006/01/02 15:04:02 PM Mon Jan")
}

func Timer(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	timeElapsed := end.Sub(start)
	fmt.Println(timeElapsed)
}

func Timer2(f func()) {
	timeStart := time.Now()
	f()
	timeElapsed := time.Since(timeStart)
	fmt.Println(timeElapsed)
}
