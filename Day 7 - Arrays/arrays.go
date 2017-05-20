package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

/*
Given an array, A, of  N integers, print A's elements in reverse order as a single line of space-separated numbers.

The first line contains an integer,  N (the size of our array).
The second line contains  N space-separated integers describing array A's elements.

4
1 4 3 2
*/

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		lines = append(lines, scanner.Text())

	}

	n, _ := strconv.Atoi(lines[0])

	line := strings.Split(lines[1], " ")

	for x := n - 1; x >= 0; x-- {

		fmt.Print(line[x] + " ")

	}

}
