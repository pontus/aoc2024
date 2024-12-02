package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		pos := 0
		last := 0
		dir := 0

		reportissafe := true
		for linescan.Scan() {

			a, err := strconv.Atoi(linescan.Text())

			if err != nil {
				panic("bad data")
			}

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

			fmt.Printf("Read %v ", a)
			last = a
			pos++
		}

		if reportissafe {
			safe++
		} else {
			unsafe++
		}

		fmt.Printf("\n")

	}

	fmt.Printf("safe %d unsafe %d \n", safe, unsafe)

}
