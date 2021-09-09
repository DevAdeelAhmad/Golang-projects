package main

import "fmt"

func main() {
	//Three ways to declare a string
	//	--> 1st -- just use the var kyword then the varName and then the type i.e. String after that "=" and then in
	//      "", value of the variable which would be a string

	var name string = "devAdeelAhmad"
	fmt.Println(name)

	//	--> 2nd -- don't put the varType after the varName and let the compiler do it's job.

	var name1 = "devAdeelAhmad"
	fmt.Println(name1)

	//	--> 3rd -- don't even use the var keyword and just put the colon(:) before assigning value i.e. before (=)
	//NOTE: This won't work outside of the functions

	name2 := "devAdeelAhmad"
	fmt.Println(name2)

	// We can also change the values of the variables afterwards

	name = "dev"
	name1 = "dev"
	name2 = "dev"

	// fmt.Println(name, name1, name2)

	//The same procedure applies for int vars
	//Three ways to create an integer var
	// ---> 1st
	var integer1 int = 30

	// ---> 2st
	var integer2 = 30

	// ---> 3rd
	integer3 := 30

	fmt.Println(integer1, integer2, integer3)

	//Useful Links
	// https://golang.org/ref/spec#Numeric_types
	// https://pkg.go.dev/std

	var numOne int16 = 400
	fmt.Println(numOne)

	//Floating Point Integers
	var scoreOne float32 = 90.5
	var scoreTwo float64 = 9899.7
	//Just printing to avoid unwanted errors
	fmt.Println(scoreOne, scoreTwo)
}
