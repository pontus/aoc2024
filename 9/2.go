package main

import (
	_ "embed"
	"fmt"
	"slices"
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

type ent struct {
	id, length int
}

func moveNext(t *[]ent) bool {
	i := len(*t) - 1

	for {
		for {

			// fmt.Printf("Seeing at %v \n", i)

			for i >= 0 && (*t)[i].id == -1 {
				i--
			}

			if i < 0 {
				return false
			}

			l := (*t)[i].length
			// fmt.Printf("Will try to move from pos %v  of length %v id  %v\n", i, l, (*t)[i].id)
			j := 0
			for j < len(*t) && j < i {
				if (*t)[j].length >= l && (*t)[j].id == -1 {
					remain := (*t)[j].length - l
					(*t)[j].id = (*t)[i].id
					(*t)[j].length = (*t)[i].length
					(*t)[i].id = -1

					if remain > 0 {
						*t = slices.Insert(*t, j+1, ent{id: -1, length: remain})
					}

					// fmt.Printf("Did move of id %v i=%v j=%v\n", (*t)[j].id, i, j)

					return true
				}
				j++
			}
			// fmt.Printf("Couldn't move %v , length %v\n", (*t)[i].id, l)
			i-- // Skip
		}
	}
	// fmt.Printf("Giving up\n")

	return false
}

func checkSum(t []ent) int {
	s := 0

	i := 0
	for _, x := range t {
		c := i
		for i < (c + x.length) {
			if x.id != -1 {
				s = s + i*x.id
			}
			i++
		}
	}

	return s
}

func main() {
	id := 0
	diskmap := make([]ent, 0)
	nexthop := false
	for _, s := range data {
		fmt.Printf("Seeing %c\n", s)
		if s < '0' || s > '9' {
			break
		}

		n := int(s - '0')

		if nexthop {
			diskmap = append(diskmap, ent{id: -1, length: n})

		} else {
			diskmap = append(diskmap, ent{id: id, length: n})
			id++
		}

		nexthop = !nexthop
	}

	fmt.Printf("Cal %v \n", diskmap)

	// for moveNext(&diskmap) {
	// 	fmt.Printf("Cal %v \n", diskmap)

	// }

	for moveNext(&diskmap) {
		// fmt.Printf("Cal %v \n", diskmap)
	}

	fmt.Printf("Cal %v \n", diskmap)
	fmt.Printf("Cal %v \n", checkSum(diskmap))

}
