package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	arguments := os.Args[1]

	switch arguments {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("value:", arguments)
	}

	value, err := strconv.Atoi(arguments)
	if err != nil {
		fmt.Println("Cannot convert to int:", arguments)
		return
	}

	switch {
	case value == 0:
		fmt.Println("Zero")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)

	}
}
