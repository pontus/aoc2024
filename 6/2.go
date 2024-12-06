package main

import (
	_ "embed"
	"encoding/json"
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
	x, y, dir int
}

func posComp(a, b pos) bool {
	if a.x == b.x && a.y == b.y && a.dir == b.dir {
		return true
	}
	return false
}

func deepCopy(src, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}

func walk(mat [][]rune, x, y int) bool {

	trail := make([]pos, 0)
	dir := 0

	for {
		if slices.ContainsFunc(trail, func(e pos) bool {
			if e.x == x && e.y == y && e.dir == dir {
				return true
			}
			return false
		}) {
			return true
		}

		trail = append(trail, pos{x: x, y: y, dir: dir})

		lookatx, lookaty := nextCoord(x, y, dir)

		if !onMap(mat, lookatx, lookaty) {
			return false
		}

		for mat[lookaty][lookatx] == '#' {
			dir = (dir + 1) % 4
			lookatx, lookaty = nextCoord(x, y, dir)

			if !onMap(mat, lookatx, lookaty) {
				return false
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

			if mat[i][j] == '^' {
				x, y = j, i
			}
		}
	}

	count := 0
	for i := range mat {
		for j := range mat[i] {

			if mat[i][j] != '^' {

				fmt.Printf("Checking %d,%d\n", j, i)
				newmat := make([][]rune, 0)
				err := deepCopy(mat, &newmat)
				if err != nil {
					panic(err)
				}

				newmat[i][j] = '#'
				if walk(newmat, x, y) {
					count++
				}
			}
		}
	}

	fmt.Printf("data %v\n\n", count)

	//fmt.Println(sum)

}
