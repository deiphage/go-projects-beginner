/* hello-world-math-very-simple
This is not meant to be in any way a detailed list of the math
capability of GO only a demonstration of some of the simple math
operators that GO uses. this may be expanded later in another folder
to show more complex math operations
*/
package main //this is an executable file

import "fmt" //importing the fmt package allows us to use Println

func main() { //Main function
	fmt.Println("20 + 3 =", 20+3) //+ to add
	fmt.Println("20 - 3 =", 20-3) //- so subtract
	fmt.Println("20 * 3 =", 20*3) //* to multiply
	fmt.Println("20 / 3 =", 20/3) //'/' to divide
	fmt.Println("20 % 3 =", 20%3) //% for mod

}
