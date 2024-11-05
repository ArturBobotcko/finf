package main

import (
	"finf/app"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: finf <filename>")
		os.Exit(1)
	}

	fileName := flag.Arg(0)
	app.PrintFileInfo(fileName)
}
