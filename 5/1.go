package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

func main() {

	lines := strings.Split(data, "\n")
	pairre := regexp.MustCompile("^(\\d+)\\|(\\d{1,3})$")
	listre := regexp.MustCompile("\\d+")

	sum := 0
	pairs := make([]pairing, 0)
	for _, s := range lines {
		pair := pairre.FindAllStringSubmatch(s, -1)

		fmt.Printf("Scan of %v returned %v\n", s, pair)
		if len(pair) > 0 {
			// Pairing
			pairs = append(pairs, pairing{before: toNum(pair[0][1]), after: toNum(pair[0][2])})
		} else {
			list := listre.FindAllString(s, -1)
			if len(list) == 0 {
				continue
			}

			fmt.Printf("Scan of %v returned %v\n", s, list)
			update := make([]int, 0)

			for i := range list {
				update = append(update, toNum(list[i]))
			}
			fmt.Printf("Scan of %v returned %v update %v \n", s, list, update)

			middle := len(update) / 2
			if updateOk(update, pairs) {
				sum += update[middle]
			}
		}

	}
	fmt.Printf("pairs: %v\n", pairs)
	fmt.Printf("sum: %v\n\n", sum)

}
