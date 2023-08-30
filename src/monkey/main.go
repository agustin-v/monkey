package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, error := user.Current()

	if error != nil {
		panic(error)
	}

	fmt.Printf("Hello is %s! This is the Monkey promgramming language\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
