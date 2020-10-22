package main

import (
	"fmt"
	"os"
)

func main() {
	for _, val := range os.Args[1:] {
		fmt.Print(val, " ")
	}
}
