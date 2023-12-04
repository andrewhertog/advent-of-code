package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	subset   []Subset
	total    Subset
	valid    bool
	minCount Subset
	power    int
}

type Subset struct {
	blue  int
	green int
	red   int
}

const (
	totalRed   = 12
	totalGreen = 13
	totalBlue  = 14
)

func main() {
	var games []Game

	unParsed, err := readInput("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range unParsed {
		g := Game{}
		g.parseGame(v)
		games = append(games, g)
	}
	var total int
	var totalPower int
	for k, v := range games {
		if v.valid {
			total += (k + 1)
		}
		v.getMinStone()
		totalPower += v.power
	}
	fmt.Println(total)

}

func readInput(fn string) ([]string, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(dat), "\n"), nil
}

func (g *Game) parseGame(game string) {
	g.valid = true
	game = strings.Split(game, ":")[1]
	substrs := strings.Split(game, ";")
	for _, v := range substrs {
		s := Subset{}
		s.parseSubset(v)
		g.incTotal(s)
		if !s.isValid() {
			g.isValid()
		}

		g.subset = append(g.subset, s)
	}
}

func (g *Game) incTotal(s Subset) {
	g.total.blue += s.blue
	g.total.green += s.green
	g.total.red += s.red
}

func (g *Game) isValid() {
	if g.valid {
		g.valid = false
	}
}

func (s *Subset) isValid() bool {
	if s.red > totalRed {
		return false
	} else if s.blue > totalBlue {
		return false
	} else if s.green > totalGreen {
		return false
	} else {
		return true
	}
}

func (s *Subset) parseSubset(subset string) {
	substrs := strings.Split(subset, ",")
	for _, v := range substrs {
		s.parseColor(v)
	}
}

func (s *Subset) parseColor(role string) {
	if isRed(role) != 0 {
		s.red = isRed(role)
	}
	if isGreen(role) != 0 {
		s.green = isGreen(role)
	}
	if isBlue(role) != 0 {
		s.blue = isBlue(role)
	}
}

func isRed(role string) int {
	re := regexp.MustCompile("red")
	res := 0
	if re.FindString(role) != "" {
		role = strings.TrimSpace(role)
		res, _ = strconv.Atoi(strings.Split(role, " ")[0])
	}
	return res
}

func isBlue(role string) int {
	re := regexp.MustCompile("blue")
	res := 0
	if re.FindString(role) != "" {
		role = strings.TrimSpace(role)
		res, _ = strconv.Atoi(strings.Split(role, " ")[0])
	}
	return res
}
func isGreen(role string) int {
	re := regexp.MustCompile("green")
	res := 0
	if re.FindString(role) != "" {
		role = strings.TrimSpace(role)
		res, _ = strconv.Atoi(strings.Split(role, " ")[0])
	}
	return res
}

func (g *Game) getMinStone() {
	for _, v := range g.subset {
		if v.blue > g.minCount.blue {
			g.minCount.blue = v.blue
		}
		if v.green > g.minCount.green {
			g.minCount.green = v.green
		}
		if v.red > g.minCount.red {
			g.minCount.red = v.red
		}
	}
	g.calcPower()
}

func (g *Game) calcPower() {
	g.power = g.minCount.blue * g.minCount.green * g.minCount.red
}
