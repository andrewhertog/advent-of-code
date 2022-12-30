package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// input, err := readInput("test_input")
	input, err := readInput("input")
	if err != nil {
		fmt.Println(err)
	}
	processedInput := processInput(input)
	total := 0
	for _, v := range processedInput {
		total += convertChartoInt(compare(v))
	}
	fmt.Println(total)
	total = 0
	for _, v := range groupings(input) {
		total += convertChartoInt(calcBadge(v))
	}
	fmt.Println(total)
}

func readInput(fn string) (string, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func processInput(input string) [][]string {
	res := strings.Split(input, "\n")
	var final [][]string
	for _, v := range res {
		final = append(final, halfString(v))
	}
	return final
}

func halfString(input string) []string {
	length := len(input) / 2
	var res []string
	first := input[:length]
	last := input[length:]
	res = append(res, first)
	res = append(res, last)
	return res
}

func compare(input []string) string {
	var res string
	for _, v := range input[0] {
		// for _, y := range input[1] {
		// 	if v == y {
		// 		if strings.Contains(res, string(v)) {
		// 			res = res + string(v)
		// 		}
		// 	}
		// }
		if strings.Contains(input[1], string(v)) {
			res = res + string(v)
		}
	}
	return res
}

func convertChartoInt(input string) int {
	const lowerOffset = 96
	const upperOffset = 64
	const alpha = 26
	r := rune(input[0])
	res := 0
	if unicode.IsUpper(r) {
		res = int(r) - upperOffset + alpha
	} else if unicode.IsLower(r) {
		res = int(r) - lowerOffset
	}
	return res
}

func groupings(input string) [][]string {
	res := strings.Split(input, "\n")
	const gS = 3
	var groups [][]string

	for k, v := range res {
		idx := k / gS
		subIdx := k % gS
		if subIdx == 0 {
			grouping := make([]string, gS)
			groups = append(groups, grouping)
		}
		groups[idx][subIdx] = v
	}
	return groups
}

func calcBadge(input []string) string {
	res := compare(input)
	res = compare([]string{res, input[2]})
	return res[0:1]
}
