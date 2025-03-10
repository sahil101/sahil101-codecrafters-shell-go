package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// var BUILTIN_COMMANDS = [4]string{"echo", "exit", "type", "pwd", "cd"}

func isShellBuiltIn(cmd string) bool {
	return cmd == "echo" || cmd == "exit" || cmd == "type" || cmd == "pwd" || cmd == "cd"
}

// handles the type command
func type_cmd(command string) {

	if isShellBuiltIn(command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else if path, err := exec.LookPath(command); err == nil { // can we write our own lookPath function?
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

func getAbsolutePath() string {
	pwd, _ := os.Getwd()
	return pwd
}

func handleChangeDirectory(path string) {

	if path == "~" {
		path = os.Getenv("HOME")
	}

	err := os.Chdir(path)

	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
		return
	}
}

func inputParser(input string) []string {
	s := strings.Trim(input, "\r\n")
	var params []string
	var current strings.Builder
	inQuote := false
	inDoubleQuote := false
	isBackSlashed := false
	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '.' {
			current.WriteByte(char)
			continue
		}

		switch char {

		case '\\':
			if !inQuote && !isBackSlashed {
				isBackSlashed = true
			} else {
				current.WriteByte(char)
				isBackSlashed = false
			}
		case '"':
			if isBackSlashed || inQuote {
				current.WriteByte(char)
				isBackSlashed = false
			} else {
				inDoubleQuote = !inDoubleQuote
			}
		case '$' | '\n':
			if inDoubleQuote && isBackSlashed {
				current.WriteByte(char)
				isBackSlashed = false
			}
		case '\'':
			if !inQuote && !inDoubleQuote && isBackSlashed {
				current.WriteByte(char)
				isBackSlashed = false
			} else if inDoubleQuote {
				if isBackSlashed {
					current.WriteByte('\\')
					isBackSlashed = false
				}
				current.WriteByte(char)
			} else {
				// Toggle qouting mode
				inQuote = !inQuote
			}
		case ' ':
			// if outside quotes, treat as a separator

			if !inQuote && !inDoubleQuote && !isBackSlashed {
				if current.Len() > 0 {
					params = append(params, current.String())
					current.Reset()
				}
			} else {
				current.WriteByte(char)
				isBackSlashed = false
			}
		default:
			if isBackSlashed && inDoubleQuote {
				current.WriteByte('\\')
				isBackSlashed = false
			}
			current.WriteByte(char)
		}
	}

	if current.Len() > 0 {
		params = append(params, current.String())
	}
	return params
}

func handleEcho(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stdout)
		return nil
	}

	for i := 0; i < len(args)-1; i++ {
		fmt.Fprintf(os.Stdout, "%s ", args[i])
	}

	fmt.Fprintln(os.Stdout, args[len(args)-1])
	return nil

}

func main() {
	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		params := inputParser(input)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		if params[0] == "echo" {
			handleEcho(params[1:])
		} else if params[0] == "type" {
			type_cmd(params[1])
		} else if params[0] == "exit" { // revisit this
			os.Exit(0)
		} else if params[0] == "pwd" {
			pwd := getAbsolutePath()
			fmt.Println(pwd)
			continue
		} else if params[0] == "cd" {
			handleChangeDirectory(params[1])
		} else if _, err := exec.LookPath(params[0]); err == nil {
			handleFileExecution(params[0], params[1:])
		} else {
			fmt.Println(strings.ToLower(params[0]) + ": command not found")
		}

	}

}
