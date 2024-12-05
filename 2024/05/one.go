package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func main() {
	pages, rules := ReadInput("input.txt")

	fmt.Println("rules", rules)
	fmt.Println("pages", pages)

	validPages := [][]int{}

	for _, page := range pages {
		if ValidateRules(page, rules) {
			validPages = append(validPages, page)
		}
	}

	fmt.Println("valid pages", validPages)

	sum := 0
	for _, page := range validPages {
		middle := len(page) / 2
		sum += page[middle]
		fmt.Println(page[middle])
	}
	fmt.Println("sum", sum)
}

func ValidateRules(p []int, rules []Rule) bool {
	for _, rule := range rules {
		ibefore := slices.Index(p, rule.Before) // prints 1
		iafter := slices.Index(p, rule.After)   // prints 1

		if ibefore == -1 || iafter == -1 {
			continue
		}
		if ibefore > iafter {
			return false
		}
	}

	return true
}

func ReadInput(path string) (p [][]int, r []Rule) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			rule := ToRule(line)
			r = append(r, rule)
		} else {
			pages := ToPages(line)
			p = append(p, pages)
		}
	}

	return p, r
}

func ToRule(input string) Rule {
	parts := strings.Split(input, "|")

	i, _ := strconv.Atoi(parts[0])
	j, _ := strconv.Atoi(parts[1])

	return Rule{Before: i, After: j}
}

func ToPages(input string) []int {
	var out []int
	parts := strings.Split(input, ",")
	for _, v := range parts {
		j, _ := strconv.Atoi(v)
		out = append(out, j)
	}
	return out
}
