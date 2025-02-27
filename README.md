# MyShell

MyShell is a simple command-line shell implemented in Go. It supports basic shell commands such as `echo`, `exit`, `type`, `pwd`, and `cd`. The shell reads user input, parses commands, and executes them accordingly.

## Features

- **Built-in Commands**: 
  - `echo`: Prints the provided arguments to the standard output.
  - `exit`: Exits the shell.
  - `type`: Displays whether the command is a built-in or its path if it's an executable.
  - `pwd`: Prints the current working directory.
  - `cd`: Changes the current directory.

- **Command Execution**: Executes external commands and captures their output.

## Installation

To install MyShell, clone the repository and build the project:

```bash
git clone https://github.com/sahil101/sahil101-codecrafters-shell-go.git
cd sahil101-codecrafters-shell-go
go build -o myshell cmd/myshell/main.go
```

## Usage

Run the shell by executing the following command:

```bash
cd /tmp/
./myshell
```

Once the shell is running, you can enter commands. For example:

```bash
$ echo Hello, World!
Hello, World!

$ pwd
/Users/yourusername/Desktop

$ cd /path/to/directory

$ type echo
echo is a shell builtin
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or features you'd like to add.

## License

This project is licensed under the MIT License.
