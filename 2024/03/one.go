package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	input := ReadInput("input.txt")

	r := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)

	matches := r.FindAllString(input, -1)

	fmt.Println(matches)

	total := 0
	for i := 0; i < len(matches); i++ {
		mul := matches[i]
		x, y := ExtractIntegers(mul)
		fmt.Println(mul, x, y)
		total += x * y
	}

	fmt.Println(total)

}

func ExtractIntegers(str string) (int, int) {
	separator := "("

	index := strings.Index(str, separator)

	if index != -1 {
		// Trim everything up to and including the separator
		result := str[index+len(separator):]
		result = result[:len(result)-1]

		parts := strings.Split(result, ",")

		// Convert the string parts to integers
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		fmt.Println("parts", parts)

		return x, y

	} else {
		log.Fatal("Separator not found!")
	}

	return 0, 0
}

func ReadInput(path string) string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}
