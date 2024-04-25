package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var minValue, maxValue float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		if i == 1 {
			minValue = n
			maxValue = n
		}
		if n < minValue {
			minValue = n
		}
		if n > maxValue {
			minValue = n
		}
	}
	fmt.Println("Min:", minValue)
	fmt.Println("Max:", maxValue)
}
