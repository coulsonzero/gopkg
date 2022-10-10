package pro

import (
	"fmt"
	"time"
	_ "unsafe" // for go:linkname
)

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
	return time.Now().Format("2006/01/02 15:04:02 PM Mon Jan")
}

func timer(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Now()
	timeElapsed := end.Sub(start)
	return timeElapsed
}

func timer2(f func()) time.Duration {
	timeStart := time.Now()
	f()
	timeElapsed := time.Since(timeStart)
	return timeElapsed
}
