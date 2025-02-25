package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isShellBuiltIn(cmd string) bool {
	return cmd == "echo" || cmd == "exit" || cmd == "type" || cmd == "pwd"
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

	// Create buffers to capture standard output and error
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		fmt.Println("Stderr:", errBuffer.String())
		return
	}
	// Print the output
	fmt.Print(outBuffer.String())
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
		} else if params[0] == "pwd" {
			pwd, _ := os.Getwd()
			fmt.Println(pwd)
			continue
		} else if _, err := exec.LookPath(params[0]); err == nil {
			handleFileExecution(params[0], params[1:])
		} else {
			fmt.Println(command + ": command not found")
		}

	}

}
