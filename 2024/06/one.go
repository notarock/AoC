package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var directions = map[string]string{
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

type Map struct {
	Map [][]string
	X   int
	Y   int
}

func main() {
	m := ReadInput("input.txt")

	records := []string{}

	for m.InBound() {
		record := m.Record()
		if slices.Index(records, record) == -1 {
			records = append(records, record)
		}
		m.Move()
	}

	fmt.Println(len(records))
}

func (m *Map) Record() string {
	return fmt.Sprintf("X: %d, Y: %d", m.X, m.Y)
}

func (m *Map) InBound() bool {
	return m.X >= 0 && m.X < len(m.Map[0]) && m.Y >= 0 && m.Y < len(m.Map)
}

func (m *Map) Move() {
	switch m.Map[m.Y][m.X] {
	case "^":
		m.GoUp()
	case "v":
		m.GoDown()
	case "<":
		m.GoLeft()
	case ">":
		m.GoRight()
	}
}

func (m *Map) GoDown() {
	if !(m.Y+1 == len(m.Map)) {
		if m.Map[m.Y+1][m.X] == "#" {
			m.Map[m.Y][m.X] = directions["v"]
			return
		}
	}

	m.Map[m.Y][m.X] = "."
	m.Y++
	if m.Y >= len(m.Map) {
		return
	}
	m.Map[m.Y][m.X] = "v"
}

func (m *Map) GoUp() {
	if !(m.Y == 0) {
		if m.Map[m.Y-1][m.X] == "#" {
			m.Map[m.Y][m.X] = directions["^"]
			return
		}
	}

	m.Map[m.Y][m.X] = "."
	m.Y--
	if m.Y < 0 {
		return
	}
	m.Map[m.Y][m.X] = "^"
}

func (m *Map) GoLeft() {
	if !(m.X == 0) {
		if m.Map[m.Y][m.X-1] == "#" {
			m.Map[m.Y][m.X] = directions["<"]
			return
		}
	}
	m.Map[m.Y][m.X] = "."
	m.X--
	if m.X < 0 {
		return
	}
	m.Map[m.Y][m.X] = "<"
}

func (m *Map) GoRight() {
	if !(m.X+1 == len(m.Map[0])) {
		if m.Map[m.Y][m.X+1] == "#" {
			m.Map[m.Y][m.X] = directions[">"]
			return
		}
	}
	m.Map[m.Y][m.X] = "."
	m.X++
	if m.X == len(m.Map[0]) {
		return
	}
	m.Map[m.Y][m.X] = ">"
}

func ReadInput(path string) Map {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	out := [][]string{}
	m := Map{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runes := strings.Split(scanner.Text(), "")
		out = append(out, runes)
	}

	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out[i]); j++ {
			if out[i][j] == "^" ||
				out[i][j] == "v" ||
				out[i][j] == "<" ||
				out[i][j] == ">" {
				m.X = j
				m.Y = i
			}
		}
	}

	m.Map = out
	return m
}
