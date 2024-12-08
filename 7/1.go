package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
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

type pairing struct {
	before, after int
}

func updateOk(update []int, pairs []pairing) bool {
	for i, u := range update {
		for _, p := range pairs {
			if p.before != u {
				// Not the one we're looking for
				continue
			}

			for x, v := range update {
				if v == p.after {
					// The pairing we want to check
					if x < i {
						return false
					}
				}
			}
		}
	}
	return true
}

func canBe(result int, vals []int, locked []int) bool {
	if len(locked) == len(vals)-1 {
		// Fully specified, evaluate result
		count := vals[0]
		for i, op := range locked {
			switch op {
			case 0:
				count = count * vals[i+1]
			case 1:
				count = count + vals[i+1]

			}
		}

		return count == result
	}

	tryMul := append(locked, 0)
	if canBe(result, vals, tryMul) {
		return true
	}
	tryAdd := append(locked, 1)
	return canBe(result, vals, tryAdd)

}

func main() {

	lines := strings.Split(data, "\n")

	cal := 0
	// sum := 0
	// pairs := make([]pairing, 0)
	for _, s := range lines {
		l := strings.Split(s, " ")
		if len(l) < 2 {
			break
		}

		// fmt.Printf("l is %v len %v\n", l, len(l))
		result := toNum(strings.Trim(l[0], ":"))
		ops := make([]int, 0)
		l = slices.Delete(l, 0, 1)

		for _, p := range l {
			ops = append(ops, toNum(p))
		}

		if canBe(result, ops, []int{}) {
			cal += result
		}

	}
	fmt.Printf("Cal %v \n", cal)

}
