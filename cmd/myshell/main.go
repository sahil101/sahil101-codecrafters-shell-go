package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// handles the type command
func type_cmd(command string) {

	if command == "echo" {
		fmt.Println("echo is a shell builtin")
	} else if command == "exit" {
		fmt.Println("exit is a shell builtin")
	} else {
		path := os.Getenv("PATH")
		fmt.Println(path)
		pathArray := strings.Split(path, ":")
		var found bool = false
		for _, p := range pathArray {
			if strings.Contains(p, command) {
				fmt.Println(command + " is " + p)
				found = true
				break
			}
		}
		if !found {
			fmt.Println(command + ": not found")
		}
	}
}

func main() {
	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.Trim(command, "\n\r")
		params := strings.Fields(command)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		switch params[0] {
		case "type":
			type_cmd(params[1])
		case "echo":
			fmt.Println(strings.Join(params[1:], " "))
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(command + ": command not found")
		}

	}

}
