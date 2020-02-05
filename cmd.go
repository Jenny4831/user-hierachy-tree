package main

import (
	"fmt"
	"strings"
)

type Command string

const (
	SETROLES        Command = "setroles"
	SETUSERS        Command = "setusers"
	GETSUBORDINATES Command = "getsubordinates"
	EXIT            Command = "exit"
	HELP            Command = "help"
)

type CMDInstruction struct {
	Command
	Instruction string
}

var CommandsInstructions = []CMDInstruction{
	{
		Command:     SETROLES,
		Instruction: `Sets roles list e.g [{ Id:1, Name: "Admin", Parent: 0 }, { Id:2, Name: "Employee", Parent: 1 }]`,
	},
	{
		Command:     SETUSERS,
		Instruction: `Sets users list e.g [{ Id:1, Name: "John", Role: 0 }, { Id:2, Name: "Tommy", Role: 1 }]`,
	},
	{
		Command:     GETSUBORDINATES,
		Instruction: `Gets subordinate of given user ID, e.g. input: 1 ; output: [{ Id:2, Name: "Tommy", Role: 1 }]`,
	},
	{
		Command:     EXIT,
		Instruction: "Exit program",
	},
}

func CoverStringToCommand(input string) Command {
	trimmed := strings.Trim(input, "\n")
	formatted := strings.ToLower(trimmed)
	return Command(formatted)
}

func PrintInstructions() {
	fmt.Println(`The following is a list of commands for this program, to execute,` +
		` hit enter after command, commands are case insensitive`)
	for idx := range CommandsInstructions {
		cmdInstruction := CommandsInstructions[idx]
		fmt.Printf("\n<%s>\n%s\n", strings.ToLower(string(cmdInstruction.Command)), cmdInstruction.Instruction)
	}
}
