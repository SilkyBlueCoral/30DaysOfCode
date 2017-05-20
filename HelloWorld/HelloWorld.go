package main

import "bufio"
import "os"
import "fmt"

func scanner() {
	// Create a Scanner to read input from stdin.
	fmt.Println("Hello, World.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { //Read in lines
		line := scanner.Text()
		if len(line) == 0 { //if line is blank, exit program
			break
		}
		fmt.Println(scanner.Text()) //Print Lines to output

	}

}


func main() {

	scanner() //Call Scanner

}
