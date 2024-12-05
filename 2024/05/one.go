package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := ReadInput("input.txt")

	totals := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if CheckCrossmas(CrossmasCheckerInput{table: input, x: j, y: i}) {
				totals++
			}
		}
	}

	fmt.Println("Total XMAS found", totals)
}

func CheckCrossmas(input CrossmasCheckerInput) bool {
	if input.table[input.y][input.x] != "A" {
		return false
	}

	if input.x == 0 ||
		input.x == len(input.table[0])-1 ||
		input.y == 0 ||
		input.y == len(input.table)-1 {
		return false
	}

	mcound := 0
	scound := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 || j == 0 {
				continue
			}
			if input.table[input.y+i][input.x+j] == "M" {
				mcound++
			}
			if input.table[input.y+i][input.x+j] == "S" {
				scound++
			}
		}
	}

	if !(mcound == 2 && scound == 2) {
		return false
	}

	// Filter out wrong patterns
	return input.table[input.y-1][input.x-1] != input.table[input.y+1][input.x+1]
}

func CheckXMAS(input XMASCheckerInput) bool {
	lenx := len(input.table[0])
	leny := len(input.table)

	// Lax position of word is out of bound on X axis
	if (input.x+3*input.directionX) >= lenx || (input.x+3*input.directionX) < 0 {
		return false
	}

	// Lay position of word is out of bound on Y ayis
	if (input.y+3*input.directionY) >= leny || (input.y+3*input.directionY) < 0 {
		return false
	}

	return input.table[input.y][input.x] == "X" &&
		input.table[input.y+(input.directionY)][input.x+(input.directionX)] == "M" &&
		input.table[input.y+2*(input.directionY)][input.x+2*(input.directionX)] == "A" &&
		input.table[input.y+3*(input.directionY)][input.x+3*(input.directionX)] == "S"
}

func ReadInput(path string) [][]string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	out := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runes := strings.Split(scanner.Text(), "")
		out = append(out, runes)
	}

	return out
}
