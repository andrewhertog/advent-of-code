package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	unParsed, err := readInput("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	values := processInput(unParsed)
	fmt.Println(sum(values))

}

func readInput(fn string) (string, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func processInput(input string) []int {
	res := strings.Split(input, "\n")
	var final []int
	for _, v := range res {
		digits := findDigits(v)
		fmt.Println(getFirst(digits))
		fmt.Println(getLast(digits))
		val, err := strconv.Atoi(fmt.Sprintf("%v%v", getFirst(digits), getLast(digits)))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		final = append(final, val)
	}
	return final
}

func findDigits(input string) []string {
	re := regexp.MustCompile("([0-9]{1})")
	return re.FindAllString(input, -1)
}

func getFirst(i []string) string {
	return i[0]
}

func getLast(i []string) string {
	len := len(i)
	return i[len-1]
}

func sum(i []int) int {
	var val int
	for _, v := range i {
		val += v
	}
	return val
}
