package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data
var data string

var width = 101
var heigth = 103
var middlew = 50
var middleh = 51

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func symmetric(machines []robot) bool {
	mat := draw(machines)

	for y := 0; y < heigth; y++ {
		for x := 0; x < middlew; x++ {
			if getAt(x, y, mat) == getAt(width-x-1, y, mat) {
				return false
			}
		}
	}

	return true
}

func treeLike(machines []robot) bool {

	mat := draw(machines)

	for y := heigth - 20; y < heigth; y++ {
		for x := middlew - 20; x < middlew; x++ {
			if !(getAt(x, y, mat) == getAt(width-1-x, y, mat)) {
				return false
			}
		}
	}

	return true
}

func highConnected(machines []robot) bool {
	mat := draw(machines)
	i := 0

	for _, m := range machines {
		x, y := m.x, m.y

		if (y > 0 && getAt(x, y-1, mat)) ||
			(y < heigth && getAt(x, y+1, mat)) ||
			(x > 0 && y > 0 && getAt(x-1, y-1, mat)) ||
			(x > 0 && y < heigth && getAt(x-1, y+1, mat)) ||
			(x < width && y > 0 && getAt(x+1, y-1, mat)) ||
			(x < width && y < heigth && getAt(x+1, y+1, mat)) ||
			(x > 0 && getAt(x-1, y, mat)) ||
			(x < width && getAt(x+1, y, mat)) {
			i++
		}
	}

	p := i * 100 / len(machines)
	// fmt.Printf("COnnected %v\n", p)
	if p > 59 {
		return true
	}
	return false
}

func draw(machines []robot) [][]bool {
	m := make([][]bool, 0)

	for y := 0; y < heigth+1; y++ {
		n := make([]bool, 0)
		for x := 0; x < width+1; x++ {
			n = append(n, getAtRaw(x, y, machines))

		}
		m = append(m, n)

	}
	return m
}

func illustrate(machines []robot) {

	mat := draw(machines)

	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			if getAt(x, y, mat) {
				fmt.Printf("*")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Printf("\n")

	}
}

func toNum(s string) int {
	// fmt.Printf("Converting %s\n", s)
	l, err := strconv.Atoi(s)

	if err != nil {
		panic("Number conversion failure")
	}
	return l
}

func getAtRaw(x, y int, machines []robot) bool {

	for _, m := range machines {
		if m.x == x && m.y == y {
			return true
		}

	}

	return false
}

func getAt(x, y int, mapdata [][]bool) bool {

	return mapdata[y][x]
}

type robot struct {
	x, y, vx, vy int
}

func parsePair(b string) (int, int) {
	s := strings.Split(strings.TrimSpace(b), ",")
	x := toNum(strings.Trim(s[0], "pv+,="))
	y := toNum(strings.Trim(s[1], "pv+,="))
	// fmt.Printf("%v**%v\n", x, y)

	return x, y
}

func oneSec(machines *[]robot) {
	for i := range *machines {
		m := (*machines)[i]
		m.x = (m.x + m.vx + width) % width
		m.y = (m.y + m.vy + heigth) % heigth

		(*machines)[i] = m
	}

	// fmt.Printf("Machines after oneSec %v\n", machines)
}

func countCuadrant(machines []robot) (int, int, int, int) {

	a, b, c, d := 0, 0, 0, 0

	for _, m := range machines {
		if m.x < middlew && m.y < middleh {
			a++
		}
		if m.x < middlew && m.y > middleh {
			c++
		}
		if m.x > middlew && m.y < middleh {
			b++
		}
		if m.x > middlew && m.y > middleh {
			d++
		}

	}
	return a, b, c, d
	// fmt.Printf("Machines after oneSec %v\n", machines)
}

func main() {

	lines := strings.Split(data, "\n")

	machines := make([]robot, 0)

	for _, s := range lines {
		l := strings.Split(s, " ")
		if len(l) < 2 {

			continue
		}

		x, y := parsePair(l[0])
		vx, vy := parsePair(l[1])
		m := robot{x: x, y: y, vx: vx, vy: vy}
		machines = append(machines, m)

	}

	c := 0
	i := 0
	for {
		oneSec(&machines)

		i++
		if highConnected(machines) {
			illustrate(machines)
			c++
			break
		}
	}

	fmt.Printf("Showed %v of %v\n", c, i)

	// a, b, c, d := countCuadrant(machines)
	// fmt.Printf("Count %v %v %v %v\n", a, b, c, d)
	// fmt.Printf("Count %v \n", a*b*c*d)

}

// fmt.Printf("Cal %v \n", cal)
