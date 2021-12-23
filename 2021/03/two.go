package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := ReadInput("test.txt")
	oxy := Oxygen(inputs)
	fmt.Println(oxy)
	co2 := CO2(inputs)
	fmt.Println(co2)
	o, _ := strconv.ParseInt(oxy, 2, 64)
	c, _ := strconv.ParseInt(co2, 2, 64)
	fmt.Println(o * c)

}

func Oxygen(inputs []string) string {
	pop := inputs
	for i := 0; i < len(inputs[0]); i++ {
		pop, _ = SplitsByPopularity(pop, i)
		if len(pop) <= 1 {
			return pop[0]
		}
	}
	return ""
}

func CO2(inputs []string) string {
	unp := inputs
	for i := 0; i < len(inputs[0]); i++ {
		_, unp = SplitsByPopularity(unp, i)
		if len(unp) <= 1 {
			return unp[0]
		}
	}
	return ""
}

func SplitsByPopularity(input []string, index int) ([]string, []string) {
	ones := []string{}
	zeros := []string{}

	for _, value := range input {
		v := strings.Split(value, "")
		if v[index] == "1" {
			ones = append(ones, value)
		} else {
			zeros = append(zeros, value)
		}
	}
	if len(ones) >= len(zeros) {
		return ones, zeros
	} else {
		return zeros, ones
	}
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
