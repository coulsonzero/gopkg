package main

import (
	"github.com/coulsonzero/gopkg/files"
)

func main() {
	// fmt.Println(gopkg.IsExists("conf"))            // true
	// fmt.Println(gopkg.IsExists("conf/.env"))       // true
	// fmt.Println(gopkg.IsExists("conf/config.ini")) // true
	// fmt.Println(gopkg.IsDir("conf"))               // true
	// fmt.Println(gopkg.IsDir("conf/config.ini"))    // false
	// fmt.Println(gopkg.IsFile("conf"))              // false
	// fmt.Println(gopkg.IsFile("conf/config.ini"))   // true

	files.WalkRemoveFiles("./", "env")

}
