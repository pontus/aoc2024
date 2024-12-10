package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed data
var data string

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toNum(s string) int {
	l, err := strconv.Atoi(s)

	if err != nil {
		panic("Number conversion failure")
	}
	return l
}

func moveNext(t *[]int) bool {
	ff := 0
	fb := len(*t) - 1

	// fmt.Printf("Looking at %v\n", *t)
	// Find first free block
	for ff < len(*t) && (*t)[ff] != -1 {
		ff++
	}

	// Find last non free block
	for fb > 0 && (*t)[fb] == -1 {
		fb--
	}

	// fmt.Printf("First free is %v, last non-free %v\n", ff, fb)

	if fb < ff {
		// Nothing to do
		return false
	}

	// fmt.Printf("Swapping \n")

	// (*t)[1] = 42
	(*t)[ff] = (*t)[fb]
	(*t)[fb] = -1
	// fmt.Printf("Looking at %v\n", *t)

	return true
}

func checkSum(t []int) int {
	s := 0
	for i, v := range t {
		if v != -1 {
			s += i * v
		}
	}
	return s
}

func main() {
	id := 0
	diskmap := make([]int, 0)
	nexthop := false
	for _, s := range data {
		fmt.Printf("Seeing %c\n", s)
		if s < '0' || s > '9' {
			break
		}

		n := int(s - '0')

		if nexthop {
			put := 0
			for put < n {
				diskmap = append(diskmap, -1)
				put++
			}

		} else {
			put := 0
			for put < n {
				diskmap = append(diskmap, id)
				put++
			}
			id++
		}

		nexthop = !nexthop
	}

	fmt.Printf("Cal %v \n", diskmap)

	for moveNext(&diskmap) {
		// fmt.Printf("Cal %v \n", diskmap)

	}
	fmt.Printf("Cal %v \n", diskmap)
	fmt.Printf("Cal %v \n", checkSum(diskmap))

}
