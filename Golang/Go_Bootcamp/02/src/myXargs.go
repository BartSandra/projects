package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)

	var args []string
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Args = append(cmd.Args, args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("An error occurred:", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}
