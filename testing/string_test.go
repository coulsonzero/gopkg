package test

import (
	"github.com/coulsonzero/gopkg"
	"testing"
)

func TestToTitle(t *testing.T) {
	org := gopkg.ToTitle("hello world!")
	if want := "Hello World!"; org != want {
		t.Errorf("error: org: %q; want: %q", org, want)
	}
}
