package main

import (
	"fmt"

	"github.com/notarock/aoc/pkg/input"
)

func main() {
	// inputString := input.ReadInputAsString("./test.txt")
	inputString := input.ReadInputAsString("./input.txt")
	current := 0
	max := 0
	for _, value := range inputString {
		fmt.Println(value)
		if value == "" {
			if max <= current {
				fmt.Println(current, " sum value")
				max = current
			}
			current = 0
		} else {
			current = current + input.StrToInt(value)
		}
	}

	fmt.Println("Maximum: ", max)
}
