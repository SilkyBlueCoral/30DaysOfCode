package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

/*
Given a string, S, of length n that is indexed from 0 to n-1,
print its even-indexed and odd-indexed characters as 2 space-separated
strings on a single line (see the Sample below for more detail).
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

	for x := 1; x <= n; x++ {
		str := lines[x]
		var str1 []byte
		var str2 []byte

		for y := 0; y < len(str); y++ {

			if y%2 == 0 {
				str1 = append(str1, str[y])

			} else {
				str2 = append(str2, str[y])
			}

		}

		fmt.Printf(string(str1[:]))
		fmt.Print(" ")
		fmt.Printf(string(str2[:]) + "\n")

	}

}
