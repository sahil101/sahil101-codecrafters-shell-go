package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

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
		case "echo":
			fmt.Println(strings.Join(params[1:], " "))
		case "exit":
			if params[1] == "0" {
				os.Exit(0)
			}

		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

		// if command == "\n" {
		// 	break
		// }

	}

}
