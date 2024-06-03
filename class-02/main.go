package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Passed args: ", os.Args[1:])
	} else {
		fmt.Println("No args.")
	}
}
