package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	start int
	end   int
}

type assignmentSection struct {
	first  assignment
	second assignment
}

func main() {
	// input, _ := readInput("test_input")
	input, _ := readInput("input")

	i := processInput(input)
	total := 0
	totalOverlap := 0
	for _, v := range *i {
		if v.isContained() {
			total += 1
		}
		if v.isOverlapped() {
			totalOverlap += 1
		}
	}
	fmt.Println(total)
	fmt.Println(totalOverlap)

}

func readInput(fn string) (string, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func processInput(input string) *[]assignmentSection {
	a := strings.Split(input, "\n")
	aS := []assignmentSection{}
	for _, v := range a {
		aS = append(aS, *processSection(v))
	}
	return &aS
}

func processSection(input string) *assignmentSection {
	i := strings.Split(input, ",")
	return createAssignmentSection(processAssignment(i[0]), processAssignment(i[1]))
}

func processAssignment(input string) *assignment {
	i := strings.Split(input, "-")
	return createAssignment(i[0], i[1])
}

func createAssignment(start string, end string) *assignment {
	s, _ := strconv.Atoi(start)
	e, _ := strconv.Atoi(end)
	return &assignment{
		start: s,
		end:   e,
	}
}

func createAssignmentSection(first *assignment, second *assignment) *assignmentSection {
	return &assignmentSection{
		first:  *first,
		second: *second,
	}
}

func (aS assignmentSection) isContained() bool {
	fa := aS.firstAssignment()
	sa := aS.secondAssignment()
	if fa.assignmentStart() <= sa.assignmentStart() && fa.assignmentEnd() >= sa.assignmentEnd() {
		return true
	}
	if fa.assignmentStart() >= sa.assignmentStart() && fa.assignmentEnd() <= sa.assignmentEnd() {
		return true
	}
	return false
}

func (aS assignmentSection) isOverlapped() bool {
	fa := aS.firstAssignment()
	sa := aS.secondAssignment()
	if fa.assignmentStart() > sa.assignmentEnd() || fa.assignmentEnd() < sa.assignmentStart() {
		return false
	}
	return true
}

func (aS assignmentSection) firstAssignment() assignment {
	return aS.first
}

func (aS assignmentSection) secondAssignment() assignment {
	return aS.second
}

func (a assignment) assignmentStart() int {
	return a.start
}

func (a assignment) assignmentEnd() int {
	return a.end
}
