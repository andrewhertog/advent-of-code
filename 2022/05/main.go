package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	count  int
	source int
	dest   int
}

type stack struct {
	items string
}

func main() {
	// input, _ := readInput("test_input")
	input, _ := readInput("input")
	stacks, moves := processInput(input)
	// fmt.Println(stacks)
	// fmt.Println(moves)
	items := ""
	for _, move := range *moves {
		for j, s := range *stacks {
			if j == move.source {
				items = s.getItems9001(move.count)
				(*stacks)[j].items = s.removeItems(move.count)

			}
		}
		for j, s := range *stacks {
			if j == move.dest {
				(*stacks)[j] = *s.addToStack(items)
			}
		}
	}
	res := ""
	for _, r := range *stacks {
		res = res + r.getItems(1)
	}
	fmt.Println(stacks)
	fmt.Println(res)

}

func readInput(fn string) (string, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func (s *stack) getItems(i int) string {
	res := ""
	for j := 0; j < i; j++ {
		res = string(s.items[j]) + res
	}
	return res
}

func (s *stack) getItems9001(i int) string {
	res := ""
	for j := 0; j < i; j++ {
		res = res + string(s.items[j])
	}
	return res
}

func (s *stack) removeItems(i int) string {
	return s.items[i:]
}

func (s *stack) addToStack(i string) *stack {
	s.items = i + s.items
	return s
}

func processInput(input string) (*[]stack, *[]move) {
	i := strings.Split(input, "\n\n")
	stacks := []stack{}
	stacks = *processStacks(i[0])
	moves := []move{}
	for _, v := range strings.Split(i[1], "\n") {
		moves = append(moves, generateMove(v))
	}
	return &stacks, &moves
}

func processStacks(input string) *[]stack {
	i := strings.Split(input, "\n")
	columns := len(i[len(i)-1]) / 4
	s := []stack{}
	for i := 0; i <= columns; i++ {
		s = append(s, stack{})
	}
	for _, row := range i[0 : len(i)-1] {
		for j := 0; j <= len(row); j++ {
			c := j / 4
			if j%4 == 2 {
				if string(row[j-1]) != " " {
					s[c] = *s[c].appendItem(string(row[j-1]))
				}
			}
		}
	}
	return &s
}

func (s *stack) appendItem(i string) *stack {
	s.items = s.items + i
	return s
}

func generateMove(input string) move {
	i := strings.Split(input, " ")
	return move{
		count:  atoi(i[1]),
		source: atoi(i[3]) - 1,
		dest:   atoi(i[5]) - 1,
	}
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
