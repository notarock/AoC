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
	first, second := ReadInput("test.txt")

	sort.Ints(first)
	sort.Ints(second)

	result := []int{}

	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			c := 0
			if first[i] == second[j] {
				c++
			}
			result = append(result, first[i]*c)
		}
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
