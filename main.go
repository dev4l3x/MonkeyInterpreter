package main

import (
	"fmt"
	"os"
	"os/user"
	"monkeyinterpreter/repl"
)

func main() {
	
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey language :)\n", user.Username)
	fmt.Printf("Start by typying some commands...\n")

	repl.Start(os.Stdin, os.Stdout)

}
