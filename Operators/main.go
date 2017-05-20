package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
/*
Task
Given an integer, , perform the following conditional actions:

If n is odd, print Weird
If n is even and in the inclusive range of  2 to 5, print Not Weird
If n is even and in the inclusive range of  6 to 20, print Weird
If n is even and greater than 20, print Not Weird
Complete the stub code provided in your editor to print whether or not  is weird.

 */

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		line=scanner.Text() // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	x,_:=strconv.ParseInt(line, 0, 64)


	if x%2!=0 || (x>=6 && x<=20){

		fmt.Println("Weird")

	}else {

		fmt.Println("Not Weird")
	}

}

