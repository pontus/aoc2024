package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed data
var data string

func handleMul(s string) int {
	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	mul := re.FindAllStringSubmatch(s, -1)

	fmt.Printf("mul(%v,%v)\n", mul[0][1], mul[0][2])
	l, _ := strconv.Atoi(mul[0][1])
	r, _ := strconv.Atoi(mul[0][2])
	return l * r
}

func nextCoord(inx, iny, dir int, mat [][]rune) (x, y int, valid bool) {

	x = inx
	y = iny
	valid = true
	switch dir {
	case 0:
		x++
	case 1:
		x++
		y++
	case 2:
		y++
	case 3:
		x--
	case 4:
		x--
		y--
	case 5:
		y--
	case 6:
		x++
		y--
	case 7:
		x--
		y++
	}

	if y < 0 || y >= len(mat) {
		valid = false
		y = 0
	}

	if x < 0 || x >= len(mat[y]) {
		valid = false
		x = 0
	}

	return
}

func XmasAt(mat [][]rune, x, y int) bool {

	if !MasAt(mat, 1, x, y) && !MasAt(mat, 4, x+2, y+2) {
		return false
	}

	if !MasAt(mat, 7, x+2, y) && !MasAt(mat, 6, x, y+2) {
		return false
	}

	// We have a Mas at x,y

	return true
}

func MasAt(mat [][]rune, dir, x, y int) bool {

	if y < 0 || y >= len(mat) {
		return false
	}

	if x < 0 || x >= len(mat[y]) {
		return false
	}

	fmt.Printf("At %d,%d checking for M, seeing %c\n", x, y, mat[y][x])

	if mat[y][x] != 'M' {
		fmt.Printf("  Giving up\n")
		return false
	}

	dx, dy, valid := nextCoord(x, y, dir, mat)

	if valid {
		fmt.Printf("At %d,%d checking for A, seeing %c\n", dx, dy, mat[dy][dx])
	} else {
		fmt.Printf("At %d,%d invalid \n", dx, dy)
	}

	if !valid || mat[dy][dx] != 'A' {
		fmt.Printf("  Giving up\n")
		return false
	}

	dx, dy, valid = nextCoord(dx, dy, dir, mat)
	if valid {
		fmt.Printf("At %d,%d checking for S, seeing %c\n", dx, dy, mat[dy][dx])
	} else {
		fmt.Printf("At %d,%d invalid \n", dx, dy)
	}

	if !valid || mat[dy][dx] != 'S' {
		fmt.Printf("  Giving up\n")
		return false
	}

	fmt.Printf("  SUCCESS!\n")

	return true
}

func main() {

	// var mat [][]rune
	// mat := make([][]rune, 12)
	// for i := range mat {
	// 	mat[i] = make([]rune, 12)
	// }
	mat := make([][]rune, 0)
	y := 0

	mat = append(mat, make([]rune, 0))

	for _, b := range data {
		//		fmt.Println(s)
		switch b {
		case 10:
			y++
			mat = append(mat, make([]rune, 0))

		default:
			mat[y] = append(mat[y], b)
		}
	}

	for i := range mat {
		for j := range mat[i] {
			fmt.Printf("%c", mat[j][i])
		}
		fmt.Printf("\n")
	}

	found := 0
	for y = range mat {
		for x := range mat[y] {
			t := XmasAt(mat, x, y)
			if t {
				found++
			}
		}
	}

	fmt.Printf("data %v\n\n", found)

	//fmt.Println(sum)

}
