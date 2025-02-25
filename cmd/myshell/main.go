package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isShellBuiltIn(cmd string) bool {
	return cmd == "echo" || cmd == "exit" || cmd == "type"
}

// handles the type command
func type_cmd(command string) {

	if isShellBuiltIn(command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else if path, err := exec.LookPath(command); err == nil {
		fmt.Printf("%s is %s\n", command, path)
	} else {
		fmt.Println(command + ": not found")
	}
}

func handleFileExecution(command string, args []string) {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing command: ", err)
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

		if params[0] == "echo" {
			fmt.Println(strings.Join(params[1:], " "))
		} else if params[0] == "type" {
			type_cmd(params[1])
		} else if params[0] == "exit" {
			os.Exit(0)
		} else if _, err := exec.LookPath(params[0]); err == nil {
			handleFileExecution(params[0], params[1:])
		} else {
			fmt.Println(command + ": command not found")
		}

	}

}
