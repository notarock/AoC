package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	first, second := ReadInput("input.txt")

	sort.Ints(first)
	sort.Ints(second)

	result := []int{}

	for i := 0; i < len(first); i++ {
		distance := first[i] - second[i]
		if distance < 0 {
			distance = -distance
		}
		result = append(result, distance)
	}

	total := 0

	for i := 0; i < len(result); i++ {
		total += result[i]
	}

	fmt.Println(total)
}

func ReadInput(path string) ([]int, []int) {
	file, err := os.Open(path)

	var first []int
	var second []int

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "   ")

		i, _ := strconv.Atoi(s[0])
		first = append(first, i)

		i, _ = strconv.Atoi(s[1])
		second = append(second, i)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return first, second
}
