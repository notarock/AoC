package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	Action string
	Value  int
}

func main() {
	inputs := ReadInput("input.txt")
	x, y := SubPosition(0, 0, inputs)
	fmt.Println(x * y)
}

func SubPosition(x int, y int, inputs []command) (int, int) {
	aim := 0
	for _, cmd := range inputs {
		switch cmd.Action {
		case "forward":
			x = x + cmd.Value
			y = y + (cmd.Value * aim)
		case "down":
			aim = aim + cmd.Value
		case "up":
			aim = aim - cmd.Value
		default:
			fmt.Println("No information available for that command?")
		}
	}
	return x, y
}

func ReadInput(path string) []command {
	file, err := os.Open(path)
	var out []command
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i, _ := strconv.Atoi(line[1])
		cmd := command{line[0], i}
		out = append(out, cmd)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return out
}
