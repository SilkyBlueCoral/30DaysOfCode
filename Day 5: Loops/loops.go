package main

import (
	"fmt"
	"strconv"
)

func main() {


	var i int
	_, _ = fmt.Scanf("%d", &i)


	for x := 1; x <= 10; x++ {

		y := i * x
		fmt.Println(strconv.Itoa(i) + " x " + strconv.Itoa(x) + " = " + strconv.Itoa(y))

	}

}
