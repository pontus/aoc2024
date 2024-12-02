package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkLine(line []int) bool {
	pos := 0
	last := 0
	dir := 0
	reportissafe := true

	for _, a := range line {
		fmt.Printf("    Read %v ", a)

		switch pos {
		case 0:
		case 1:
			if a > last {
				dir = 1
			} else if a < last {
				dir = -1
			} else {
				reportissafe = false
			}
			if abs(a-last) > 3 || a == last {
				reportissafe = false
			}
		default:
			if dir == 1 && a <= last {
				reportissafe = false
			}
			if dir == -1 && a >= last {
				reportissafe = false
			}

			if abs(a-last) > 3 || a == last {
				reportissafe = false
			}
		}

		last = a
		pos++
	}

	fmt.Printf("%v\n", reportissafe)
	return reportissafe
}

func main() {
	f, _ := os.Open("data")
	defer f.Close()

	safe := 0
	unsafe := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		l := bytes.NewBufferString(scanner.Text())

		linescan := bufio.NewScanner(l)
		linescan.Split(bufio.ScanWords)

		line := make([]int, 0)
		for linescan.Scan() {
			a, err := strconv.Atoi(linescan.Text())
			if err != nil {
				panic("bad data")
			}

			line = append(line, a)
		}

		fmt.Printf("Line %v\n", line)

		reportissafe := checkLine(line)

		if !reportissafe {
			for i, _ := range line {

				testline := make([]int, len(line))
				copy(testline, line)

				testline = slices.Delete(testline, i, i+1)

				fmt.Printf("  Checking %v\n", testline)
				reportissafe = checkLine(testline)

				if reportissafe {
					break
				}
			}

		}

		if reportissafe {
			safe++
		} else {
			unsafe++
		}

		fmt.Printf("%v \n", reportissafe)

	}

	fmt.Printf("safe %d unsafe %d \n", safe, unsafe)

}
