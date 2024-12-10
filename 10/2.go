package main

import (
	_ "embed"
	"fmt"
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

func onMap(mat [][]int, x, y int) bool {
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

func walk(mat [][]int, x, y int) []pos {

	v := mat[y][x]
	fmt.Printf("Walk called for %d,%d - %d\n", x, y, v)

	if v == 9 {
		l := pos{}
		l.x = x
		l.y = y
		return []pos{l}
	}

	poses := make([]pos, 0)
	for dir := range []int{0, 1, 2, 3} {
		dx, dy := nextCoord(x, y, dir)
		if onMap(mat, dx, dy) && mat[dy][dx] == v+1 {
			poses = append(poses, walk(mat, dx, dy)...)
		}
	}
	return poses
}

func main() {

	mat := make([][]int, 0)
	y := 0

	allScore := 0
	mat = append(mat, make([]int, 0))

	for _, b := range data {
		switch b {
		case 10:
			y++
			mat = append(mat, make([]int, 0))

		default:
			mat[y] = append(mat[y], int(b-'0'))
		}
	}

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				fmt.Printf("Starting from %d,%d ", j, i)
				poses := walk(mat, j, i)
				// slices.SortFunc(poses, posSort)
				// poses = slices.CompactFunc(poses, posComp)
				allScore += len(poses)

			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("data %v\n\n", allScore)

	// fmt.Printf("data %v\n\n", pos)

	// fmt.Printf("data %v\n\n", len(pos))

	fmt.Println(allScore)

}
