package gopkg

import (
	"fmt"
	"log"
)

const prefix = "| gopkg | "

func logger(e any) {
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(fmt.Sprintf("\n%c[0;34;34m%s%c[0m", 0x1B, prefix, 0x1B))
	log.Fatal(e)
}
