package main

import (
	"github.com/Kalimaha/toyrobot-go/pkg/toyrobot"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Instructions())
	} else {
		filepath := args[1]
		_, _ = toyrobot.Play(filepath, toyrobot.INFO)
	}
}

func Instructions() string {
	return `
Please provide the absolute path to the file containing the instruction for the Toy Robot
e.g. go run ./toyrobot /tmp/example.txt
`
}
