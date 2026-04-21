package main

import (
	"fmt"
	"os"
	"goland/01-simple-example/hello"
	"goland/01-simple-example/greetings"
)

func main () {
		fmt.Println(hello.Say(os.Args[1:]))
		fmt.Println(greetings.Morning())
}
