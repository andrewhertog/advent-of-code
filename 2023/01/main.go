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
		if len(v) != 0 {
			digits := convStringtoInt(v)
			number := fmt.Sprintf("%v%v", getFirst(digits), getLast(digits))
			fmt.Printf("%v - %v - %v\n", v, digits, number)
			val, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			final = append(final, val)
		}
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

func convStringtoInt(input string) []string {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ints := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var new []string
	for k, _ := range input {
		for kk, vv := range numbers {
			if strings.Index(input[k:len(input)], vv) == 0 {
				new = append(new, ints[kk])
			}
		}
		for kk, vv := range ints {
			if strings.Index(input[k:len(input)], vv) == 0 {
				new = append(new, ints[kk])
			}
		}
	}
	return new
}
