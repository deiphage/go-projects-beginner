// hello-world-variables
package main // this is an executable file

import (
	"fmt" // allows us to use the Println function to print numbers
)

func main() { // main function
	var age int = 25            //creating an int with 25 as the value named age
	var favNum float64 = 1.5496 //creating a float64 with value 1.5496 called favNum
	randNum := 1                //The := syntax is shorthand for declaring and initializing a variable
	var v1, v2 int = 1, 2       // declaring and initializing two variables at once

	fmt.Println(age, favNum, randNum, v1, v2) // printing all the things
}
