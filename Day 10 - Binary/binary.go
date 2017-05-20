package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ones(binStr string) (int) {

	if binStr == "1" {
		return 1
	}

	vars := strings.Split(binStr, "")
	total := 0
	y := 0
	for x := 0; x < len(vars); x++ {

		fmt.Println(vars[x])
		if vars[x] == "1" {
			y += 1
		} else {
			if y > total {

				total = y
			}
			y = 0
		}

		if y > total {

			total = y
		}

	}

	return total

}

/*
Given a base-10 integer, n, convert it to binary (base-2).
Then find and print the base-10 integer denoting the maximum number of consecutive 1's in n's binary representation.
 */

func main() {

	var i int64
	_, _ = fmt.Scanf("%d", &i)

	i_bin := strconv.FormatInt(i, 2)

	fmt.Println(strconv.Itoa(ones(i_bin)))

}
