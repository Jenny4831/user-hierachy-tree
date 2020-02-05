package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		input, inputErr := reader.ReadString('\n')
		if inputErr != nil {
			fmt.Fprintln(os.Stderr, inputErr)
		}

		command := CoverStringToCommand(input)

		switch command {
		case SETROLES:
			os.Exit(0)
		case SETUSERS:
			os.Exit(0)
		case GETSUBORDINATES:
			os.Exit(0)
		case HELP:
			PrintInstructions()
		case EXIT:
			os.Exit(0)
		default:
			PrintInstructions()
		}
	}
}
