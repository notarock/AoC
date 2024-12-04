package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type XMASCheckerInput struct {
	table      [][]string
	x          int
	y          int
	directionX int
	directionY int
}

func main() {
	input := ReadInput("input.txt")

	totals := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			founds := []bool{}

			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == 0 && l == 0 {
						continue
					}

					found := CheckXMAS(XMASCheckerInput{input, j, i, k, l})
					if found {
						fmt.Println("Found at", j, i, k, l)
					}
					founds = append(founds, found)
				}
			}

			for _, found := range founds {
				if found {
					totals++
				}
			}
		}
	}

	fmt.Println("Total XMAS found", totals)
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
