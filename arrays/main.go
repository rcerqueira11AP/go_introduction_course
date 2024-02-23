package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myArray [6]int
	fmt.Println(myArray)

	myArrayVar1 := [3]string{"a", "b", "c"}
	fmt.Println(myArrayVar1)

	myArray[1] = 2
	myArray[2] = 5
	myArray[3] = 9
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArray, unsafe.Sizeof(myArray), unsafe.Sizeof(myArray)*8)

	var slice1 []int
	fmt.Printf("size: %d, value: %v\n", len(slice1), slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Printf("size: %d, value: %v\n", len(slice1), slice1)

	mySlice := myArray[2:4] // no se incluye el ultimo valor [4] no se toma
	fmt.Println(mySlice)
	fmt.Printf("size: %d, value: %v\n", len(mySlice), mySlice)

	fmt.Println(&myArray[2])
	fmt.Println(&mySlice[0])

	slice := make([]int, 3)
	fmt.Printf("size: %d, value: %v\n", len(slice), slice)
}
