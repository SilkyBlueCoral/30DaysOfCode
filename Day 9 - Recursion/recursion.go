package main

import (
	"fmt"
	"strconv"
)

func factorial(num int) (int) {

	if num == 1 {
		return 1
	}
	return num * factorial(num-1)

}

func main() {

	var i int
	_, _ = fmt.Scanf("%d", &i)

	r := factorial(i)

	fmt.Print(strconv.Itoa(r))

}
