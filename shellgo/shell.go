package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func exec_from_input(input string) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	if args[0] == "exit" {
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout

	return cmd.Run()

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$$> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		error := exec_from_input(input)
		if error != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
