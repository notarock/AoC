package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInputAsInt(path string) (out []int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, StrToInt(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return out
}

func ReadInputAsString(path string) (out []string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return out
}

func StrToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println(err)
	}

	return out
}
