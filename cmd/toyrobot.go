package main

import (
	"github.com/Kalimaha/toyrobot-go/pkg/toyrobot"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println()
		fmt.Println("Please provide the absolute path to the file containing the instruction for the Toy Robot")
		fmt.Println("e.g. go run ./toyrobot /tmp/example.txt")
		fmt.Println()
	} else {
		filepath := args[1]
		toyrobot.Play(filepath, toyrobot.INFO)
	}
}
