package main

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed data
var data string

func nextCoord(inx, iny, dir int) (outx, outy int) {
	switch dir {
	case 0: // Up
		outx = inx
		outy = iny - 1
	case 1: // Right
		outx = inx + 1
		outy = iny
	case 2: // Down
		outx = inx
		outy = iny + 1
	case 3: // Left
		outx = inx - 1
		outy = iny

	}

	return
}

func onMap(mat [][]rune, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if y >= len(mat) {
		return false
	}

	if x >= len(mat[y]) {
		return false
	}

	return true
}

type pos struct {
	x, y int
}

func posSort(a, b pos) int {
	if a.x == b.x && a.y == b.y {
		return 0
	}

	if a.x == b.x {
		if a.y < b.y {
			return -1
		}
		return 1
	}

	if a.x < b.x {
		return -1
	}

	return 1
}

func posComp(a, b pos) bool {
	if a.x == b.x && a.y == b.y {
		return true
	}
	return false
}

func walk(mat [][]rune, x, y int) []pos {

	trail := make([]pos, 0)
	dir := 0

	for {
		trail = append(trail, pos{x: x, y: y})

		lookatx, lookaty := nextCoord(x, y, dir)

		if !onMap(mat, lookatx, lookaty) {
			return trail
		}

		for mat[lookaty][lookatx] == '#' {
			dir = (dir + 1) % 4
			lookatx, lookaty = nextCoord(x, y, dir)

			if !onMap(mat, lookatx, lookaty) {
				return trail
			}
		}

		x, y = lookatx, lookaty
	}
}

func main() {

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

	x, y := 0, 0
	for i := range mat {
		for j := range mat[i] {
			fmt.Printf("%c", mat[i][j])

			if mat[i][j] == '^' {
				x, y = j, i
			}
		}
		fmt.Printf("\n")
	}

	pos := walk(mat, x, y)
	fmt.Printf("data %v\n\n", pos)

	slices.SortFunc(pos, posSort)
	pos = slices.CompactFunc(pos, posComp)

	fmt.Printf("data %v\n\n", pos)

	fmt.Printf("data %v\n\n", len(pos))

	//fmt.Println(sum)

}
