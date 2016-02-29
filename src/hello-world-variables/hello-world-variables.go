// hello-world-variables
package main // this is an executable file

import (
	"fmt" // allows us to use the Println function to print numbers
)

func main() { // main function
	var age int = 25            //creating an int with 25 as the value named age
	var favNum float64 = 1.5496 //creating a float64 with value 1.5496 called favNum
	randNum := 1

	fmt.Println(age, favNum, randNum)
}
