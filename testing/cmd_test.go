package main

import (
	"github.com/coulsonzero/gopkg/pro"
	"testing"
)

func TestCmd(t *testing.T) {
	if err := pro.Command("ls"); err != nil {
		t.Error("error: exec cmd command !")
	}
}
