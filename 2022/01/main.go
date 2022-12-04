package main

import (
	"fmt"

	"github.com/notarock/aoc/pkg/input"
)

type TopCalories struct {
	one   int
	two   int
	three int
}

func (t TopCalories) Total() int {
	return t.one + t.three + t.two
}

func (t *TopCalories) Update(current int) {
	if current > t.one {
		t.three = t.two
		t.two = t.one
		t.one = current
		return
	}

	if current > t.two {
		t.three = t.two
		t.two = current
		return
	}

	if current > t.three {
		t.three = current
		return
	}
}

func main() {
	inputString := input.ReadInputAsString("./input.txt")
	// inputString := input.ReadInputAsString("./input.txt")
	current := 0
	var max TopCalories
	for i, value := range inputString {
		if value == "" {
			fmt.Println("Checking total of ", current)
			max.Update(current)
			current = 0
		} else {
			current = current + input.StrToInt(value)
		}

		if i == len(inputString)-1 {
			fmt.Println("Checking total of ", current)
			max.Update(current)
			current = 0
		}
	}

	fmt.Println("Maximum: ", max)
	fmt.Println("Total: ", max.Total())
}
