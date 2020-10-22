package main

import (
	"os"
	"fmt"
)

func main() {
	for _, val := range os.Args[1:] {
		fmt.Print(val, " ")
	}
}
