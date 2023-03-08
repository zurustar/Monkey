package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)


func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! this is the Monky!\n", user.Username)
	fmt.Printf("type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}


