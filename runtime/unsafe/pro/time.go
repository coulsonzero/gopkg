package pro

import (
	"fmt"
	"time"
	_ "unsafe" // for go:linkname
)

const TimeFormat string = "2006-01-02 15:04:05" // 24h 制
const TimeHalf string = "2006-01-02 03:04:05"   // 12h 制

//go:linkname timeNow github.com/coulsonzero/gopkg/pro.TimeNow
func timeNow() {
	t := time.Now()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}

//go:linkname timeFormat github.com/coulsonzero/gopkg/pro.TimeFormat
func timeFormat() string {
	// return time.Now().Format("2006/01/02 15:04:02 PM Mon Jan")
	return time.Now().Format(TimeFormat)
}

func timeElapsed(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Now()
	elapsed := end.Sub(start)
	return elapsed
}

func timeElapsed2(f func()) time.Duration {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return elapsed
}
