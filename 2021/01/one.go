package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := ReadInput("input-one.txt")
	out := CountDepthIncreases(inputs)
	fmt.Println(out)
}

func CountDepthIncreases(inputs []int) int {
	total := 0
	last := int(^uint(0) >> 1)

	for _, value := range inputs {
		if value > last {
			total++
		}
		last = value
	}
	return total
}

func ReadInput(path string) []int {
	file, err := os.Open(path)
	var out []int
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		out = append(out, i)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return out
}
