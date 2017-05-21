package main

import (
	"fmt"
)

func scan6x6() [6][6]int {

	var j [6][6]int

	for x := 0; x < 6; x++ { //Read in lines

		var i [6]int
		_, _ = fmt.Scanln(&i[0], &i[1], &i[2], &i[3], &i[4], &i[5])
		j[x] = i

	}

	return j
}

func add_hourglass(k [][]int) int {

	sum := k[0][0] + k[0][1] + k[0][2] + k[1][1] + k[2][0] + k[2][1] + k[2][2]

	return sum
}

func main() {

	j := scan6x6()
	//fmt.Println(j)
	/*
	Lowest amount can be
	{-9 -9 -9}
	{-9 -9 -9}
	{-9 -9 -9}
	Highest amount can be
	{9 9 9}
	{9 9 9}
	{9 9 9}
	*/
	total := -81
	for x := 0; x < 4; x++ {
		var z [][]int
		for y := 0; y < 4; y++ {

			/*
			j[x][y], j[x][y+1], j[x][y+2]
			j[x+1][y], j[x+1][y+1], j[x+1][y+2]
			j[x+2][y], j[x+2][y+1], j[x+2][y+2]
			*/

			z = [][]int{(j[x])[y:y+3], (j[x+1])[y:y+3], (j[x+2])[y:y+3]}
			sum := add_hourglass(z)
			if sum > total {
				total = sum

			}

		}

	}

	fmt.Println(total)

}
