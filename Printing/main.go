package main

import "fmt"

func main() {

	var name string = "Adeel Ahmad"
	var age int8 = 19
	var passion string = "UI/UX Designer and Developer"
	//Print Statements
	// ---> Print line
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("my name is ", name, "and my age is ", age)

	// Format Specifier documentation https://www.educative.io/edpresso/how-to-use-the-printf-function-in-golang
	// https://pkg.go.dev/fmt
	// ---> Print formatted
	fmt.Printf("My :-\n name is %v,\n age is %v,\n passion is %v \n", name, age, passion)
	fmt.Printf("My :-\n name is %q,\n age is %q,\n passion is %q \n", name, age, passion)
	//Get type of the variable by this
	fmt.Printf("The variables name, age, and passion are of type %T, %T, and %T", name, age, passion)

	//Sprintf to get the string that is printed in the terminal and to store it in a vaiable e.g.
	var str = fmt.Sprintf("My :-\n name is %q,\n age is %q,\n passion is %q \n", name, age, passion)
	fmt.Println(str)
}
