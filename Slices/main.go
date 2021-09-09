package main

import "fmt"

func main() {
	// The only difference b/w arrays and slices is that there is no definite
	// indexes in slices i.e. you can continue without setting the size of the slices
	// and also there are some other features in slices which are below.

	// ---> Initializing
	var arr = []int{5, 6}
	fmt.Println(arr)
	// ---> Appending
	arr = append(arr, 56)
	fmt.Println(arr, len(arr))

	//To copy the content of an array we use the range method.First declaring array e.g:-
	//NOTE : if there is no index num after or before the colon then it takes all the after/before till start/end reaches
	rangeOne := arr[1:]
	rangeTwo := arr[:]
	fmt.Println(rangeOne , rangeTwo)
	// You can also use th appending in the range as they are also slices
	rangeOne = append(rangeOne , 69)
	rangeTwo = append(rangeTwo , 69)
	fmt.Println(rangeOne , rangeTwo)

}
