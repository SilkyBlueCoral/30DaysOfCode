package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for x := 1; x <= 10; x++ {

		y := i * x;
		fmt.Println(strconv.Itoa(i) + " x " + strconv.Itoa(x) + " = " + strconv.Itoa(y))

	}

}
