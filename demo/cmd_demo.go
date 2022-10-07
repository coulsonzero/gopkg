package main

import (
	"github.com/coulsonzero/gopkg"
	"log"
)

func main() {
	if err := gopkg.CmdExec("wc -l LICENSE"); err != nil {
		log.Fatal(err)
	}
	log.Fatal("success: exec cmd command.")
}
