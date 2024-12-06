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

	invalidPages := [][]int{}
	validatedPages := [][]int{}

	for _, page := range pages {
		brokenRule, valid := ValidateRules(page, rules)
		if !valid {
			invalidPages = append(invalidPages, page)
		}
		fmt.Println("Invalid", page)
		fmt.Println("broken rule was", brokenRule)
	}

	// Fix all the invalid pages
	for _, page := range invalidPages {
		valid := false

		for !valid {
			rule, v := ValidateRules(page, rules)
			if !v {
				page = FixPage(page, rule)
				valid, page = DumbassBruteforceValidater(page, rules)
				_, v = ValidateRules(page, rules)
			}
			fmt.Println("Fixed rule", rule, "on pages", page)

			valid = v
		}

		validatedPages = append(validatedPages, page)
	}

	sum := 0
	for _, page := range validatedPages {
		middle := len(page) / 2
		sum += page[middle]
		fmt.Println(page[middle])
	}
	fmt.Println("sum", sum)
}

func FixPage(p []int, rules Rule) []int {
	// Find the two numbers that are out of order
	ibefore := slices.Index(p, rules.Before)
	iafter := slices.Index(p, rules.After)

	// Swap them
	p[ibefore], p[iafter] = p[iafter], p[ibefore]

	return p
}

func DumbassBruteforceValidater(p []int, rules []Rule) (bool, []int) {
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			cp := make([]int, len(p))
			copy(cp, p)
			cp[i], cp[j] = cp[j], cp[i]
			_, valid := ValidateRules(cp, rules)
			if valid {
				return true, cp
			}
		}
	}
	return false, p
}

func ValidateRules(p []int, rules []Rule) (Rule, bool) {
	for _, rule := range rules {
		ibefore := slices.Index(p, rule.Before) // prints 1
		iafter := slices.Index(p, rule.After)   // prints 1

		if ibefore == -1 || iafter == -1 {
			continue
		}
		if ibefore > iafter {
			return rule, false
		}
	}

	return Rule{}, true
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
