package main

import (
	"flag"

	"github.com/muroon/zmlog"
)

func main() {
	fpath := flag.String("f", "", "target file path")
	flag.Parse()

	err := zmlog.ParseAndGenerate(*fpath)
	if err != nil {
		panic(err)
	}
}
