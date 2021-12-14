package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := ReadInput("test.txt")
	counts := CountBits(inputs)
	gamma, epsilon := Diagnostic(counts, len(inputs))
	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Println(gamma * epsilon)
}

func Diagnostic(counts [12]int, l int) (int, int) {
	threshold := l / 2
	g := ""
	e := ""

	for _, value := range counts {
		if value > threshold {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	ei, _ := strconv.ParseInt(e, 2, 64)
	gi, _ := strconv.ParseInt(g, 2, 64)
	return int(gi), int(ei)
}

func CountBits(input []string) [12]int {
	out := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, value := range input {
		bits := strings.Split(value, "")
		for i := 0; i < 12; i++ {
			if bits[i] == "1" {
				out[i] = out[i] + 1
			}
		}
	}
	return out
}

func ReadInput(path string) []string {
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

	return out
}
