package main

import (
	"fmt"

	"github.com/pko89403/panicRec"
)

func main() {
	fmt.Println("Hello World!")

	panicRec.FileOpen("noNamedFile")
}
