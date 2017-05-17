package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		lines = append(lines, scanner.Text())

	}

	i2, _ := strconv.ParseUint(lines[0], 0, 64)
	i2 = i2 + i

	d2, _ := strconv.ParseFloat(lines[1], 64)
	d2 += d

	s2 := s + lines[2]

	//Print Lines to output

	fmt.Println(i2)
	fmt.Println(strconv.FormatFloat(d2, 'f', 1, 64))
	fmt.Println(s2)
}
