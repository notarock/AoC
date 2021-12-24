package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	start [2]int
	end   [2]int
}

type diagram [][]int

func main() {
	coords := ReadInput("test.txt")
	diag := BuildDiagram(coords)
	diag.Print()
	fmt.Println(diag.CountOverlapping())
}

func (d diagram) Print() {
	for _, x := range d {
		fmt.Println(x)
	}
}

func (d diagram) CountOverlapping() int {
	count := 0
	for _, x := range d {
		for _, val := range x {
			if val > 1 {
				count++
			}
		}
	}
	return count
}

func BuildDiagram(inputs []line) diagram {
	d := make([][]int, 1000)
	for i := range d {
		d[i] = make([]int, 1000)
	}

	for _, value := range inputs {
		switch {
		// x
		case value.start[0] < value.end[0] && value.start[1] == value.end[1]:
			for i := value.start[0]; i <= value.end[0]; i++ {
				x := i
				y := value.end[1]
				d[y][x]++
			}
		case value.end[0] < value.start[0] && value.start[1] == value.end[1]:
			for i := value.end[0]; i <= value.start[0]; i++ {
				x := i
				y := value.end[1]
				d[y][x]++
			}
		// y
		case value.start[1] < value.end[1] && value.start[0] == value.end[0]:
			for i := value.start[1]; i <= value.end[1]; i++ {
				x := value.end[0]
				y := i
				d[y][x]++
			}
		case value.end[1] < value.start[1] && value.start[0] == value.end[0]:
			for i := value.end[1]; i <= value.start[1]; i++ {
				x := value.end[0]
				y := i
				d[y][x]++
			}
		}
	}
	return d
}

func ReadInput(path string) []line {
	file, err := os.Open(path)
	var out []string
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

	lines := []line{}
	for _, l := range out {
		coords := strings.Split(l, " -> ")
		s := strings.Split(coords[0], ",")
		e := strings.Split(coords[1], ",")

		line := line{
			start: strToInt(s),
			end:   strToInt(e),
		}
		lines = append(lines, line)
	}

	return lines
}

func strToInt(in []string) [2]int {
	out := [2]int{}
	out[0], _ = strconv.Atoi(in[0])
	out[1], _ = strconv.Atoi(in[1])

	return out
}
