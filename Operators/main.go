package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)


/*
There are  lines of numeric input:
The first line has a double,  mealCost (the cost of the meal before tax and tip).
The second line has an integer,  tipPercent(the percentage of mealCost being added as tip).
The third line has an integer,  taxPercent(the percentage of mealCost being added as tax).
*/
func main() {


	scanner := bufio.NewScanner(os.Stdin)
	var lines []string


	for scanner.Scan() {
		if len(scanner.Text())==0{
			break
		}
		lines = append(lines, scanner.Text())

	}
	mealCost,_ :=strconv.ParseFloat(lines[0], 64)
	tipPercent,_ :=strconv.ParseFloat(lines[1],64)
	taxPercent,_ :=strconv.ParseFloat(lines[2],64)

	tip:=mealCost*(tipPercent/100)
	tax:=mealCost*(taxPercent/100)

	cost:=mealCost+tip+tax

	s:="The total meal cost is "+strconv.FormatFloat(cost,'f',0,64)+" dollars."

	fmt.Println(s)
	//Print Lines to output
}