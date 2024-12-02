package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput("input.txt")

	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
	safety := []bool{}

	for i := 0; i < len(input); i++ {
		safep := IsSafe(input[i])
		safety = append(safety, safep)
	}
	fmt.Println(safety)

	c := 0
	for i := 0; i < len(safety); i++ {
		if safety[i] {
			c++
		}
	}
	fmt.Println(c)
}

func IsSafe(levels []int) bool {
	return (SafeIncrease(levels) || SafeDecrease(levels)) && SafeInstability(levels)
}

func SafeIncrease(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i-1] > levels[i] {
			return false
		}
		if (levels[i-1] - levels[i]) < -3 {
			return false
		}
	}
	return true
}

func SafeDecrease(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i-1] < levels[i] {
			return false
		}
		if (levels[i-1] - levels[i]) > 3 {
			return false
		}

	}
	return true
}

func SafeInstability(levels []int) bool {
	increasing := 0

	for i := 1; i < len(levels); i++ {
		if levels[i-1] < levels[i] {
			if increasing == -1 {
				return false
			}
			increasing = 1
		}

		if levels[i-1] > levels[i] {
			if increasing == 1 {
				return false
			}
			increasing = -1
		}

		if levels[i-1] == levels[i] {
			return false
		}
	}

	return true
}

func ReadInput(path string) [][]int {
	file, err := os.Open(path)

	var first [][]int

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		line := []int{}
		for i := 0; i < len(s); i++ {
			j, _ := strconv.Atoi(s[i])
			line = append(line, j)
		}
		first = append(first, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return first
}
