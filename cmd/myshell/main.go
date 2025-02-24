package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		switch command {
		case "exit 0\n":
			os.Exit(0)
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

		// if command == "\n" {
		// 	break
		// }

	}

}
