package main

import "fmt"

func main() {
	//Three ways to create an array
	// ---> (1).
	var arr [6]int = [6]int{56, 56, 56}
	// ---> (2).
	var arr1 = [3]int{56, 56, 56}
	// ---> (3).
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr, len(arr))
	fmt.Println(arr1, len(arr1))
	fmt.Println(arr2, len(arr2))

	//change any element of the array if you know the index.
	arr2[0] = 4
	fmt.Println(arr2, len(arr2))

}
