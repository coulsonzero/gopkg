package testing

import (
	"github.com/coulsonzero/gopkg/pro"
	"testing"
)

func TestToTitle(t *testing.T) {
	org := pro.ToTitle("hello world!")
	if want := "Hello World!"; org != want {
		t.Errorf("error: org: %q; want: %q", org, want)
	}
}
