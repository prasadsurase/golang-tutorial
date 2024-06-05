package main

import (
	"fmt"
	"hello/class-02/hello"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Passed args: ", os.Args[1:])
		fmt.Println(hello.Say(os.Args[1:]))
	} else {
		fmt.Println("No args.")
		fmt.Println(hello.Say([]string{"world"}))
	}
}
