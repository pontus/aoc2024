package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func count(needle int, haystack []int) int {
	c := 0
	for _, v := range haystack {
		if v == needle {
			c++
		}
	}
	return c
}

func main() {
	f, _ := os.Open("data")
	defer f.Close()

	left := make([]int, 00)
	right := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var a, b int
		l := scanner.Text()
		//fmt.Printf("a=%d b=%d\n", a, b)
		_, _ = fmt.Sscanf(l, "%d   %d", &a, &b)

		left = append(left, a)
		right = append(right, b)
	}

	slices.Sort(left)
	slices.Sort(right)

	sim := 0

	for i, _ := range left {
		sim += left[i] * count(left[i], right)
	}

	fmt.Printf("left=%v right=%v %d\n", left, right, sim)

}
