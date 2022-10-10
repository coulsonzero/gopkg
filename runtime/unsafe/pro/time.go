package runtime

import (
	"fmt"
	"time"
	_ "unsafe" // for go:linkname
)

//go:linkname timeNow github.com/coulsonzero/gopkg/pro.Now
func timeNow() {
	t := time.Now()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}
