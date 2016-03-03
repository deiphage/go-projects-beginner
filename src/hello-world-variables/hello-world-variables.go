/* hello-world-variables
uint8  : 0 to 255
uint16 : 0 to 65535
uint32 : 0 to 4294967295
uint64 : 0 to 18446744073709551615
int8   : -128 to 127
int16  : -32768 to 32767
int32  : -2147483648 to 2147483647
int64  : -9223372036854775808 to 9223372036854775807

int, uint, and uintptr are 32 bits on 32 bit
 system normally and 64 on a 64 bit system.
*/
package main // this is an executable file

import (
	"fmt" // allows us to use the Println function to print numbers
)

func main() { // main function
	//numbers
	var (
		varA uint8   = 28        //unsigned 8 bit integers (0 to 255)
		varB uint16  = 785       //unsigned 16 bit integers
		varC uint32  = 4873      //unsigned 32 bit integers
		varD uint64  = 234032    //unsigned 64 bit integers
		varE int8    = 80        //signed 8 bit integers (-128 to 127)
		varF int16   = 456       //signed 16 bit integers
		varG int32   = 654564    //signed 32 bit integers
		varH int64   = 654654654 //signed 64 bit integers
		varI int     = 32
		varJ float32 = 32.0156     //32 bit float
		varK float64 = 65.31321654 //64 bit float
	)
	/* just need to do something with the variables and printing them
	seemed best 	*/
	fmt.Println(varA, varB, varC, varD, varE, varF, varG, varH, varI, varJ, varK)
	//declaring some more variables and printing things again.
	const age int = 25                        //creating a constant int with 25 as the value named age - value never changes
	var favNum float64 = 1.5496               //creating a float64 with value 1.5496 called favNum
	randNum := 1                              //The := syntax is shorthand for declaring and initializing a variable
	var v1, v2 int = 1, 2                     // declaring and initializing two variables at once
	fmt.Println(age, favNum, randNum, v1, v2) // printing all the things

	/*going to put some strings here soon*/
	/*going to put some custom types here soon*/
}
