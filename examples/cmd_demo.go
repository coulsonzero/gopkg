package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
	"log"
)

func main() {
	err := pro.CmdExec("ls")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("success exec cmd")
	}

}

func DemoCmdOutput() {
	output, err := pro.CmdOutput("ls")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}