package main

import "fmt"

func main() {

	var A byte = 'A'
	var a byte = 'a'

	var R byte = 82 // se le pasa el codigo ascii que lo representa
	var s byte = 115
	vector := []byte{65, 97, 82, 115}

	fmt.Println()
	fmt.Println("Value in ASCII Code: ")
	fmt.Println(A)
	fmt.Println(a)
	fmt.Println(R)
	fmt.Println(s)
	fmt.Println(vector)

	// para saber que valores respresentan en el codigo ascii se le pasa la palabra string

	fmt.Println()
	fmt.Println("Value in String: ")
	fmt.Println(string(A))
	fmt.Println(string(a))
	fmt.Println(string(R))
	fmt.Println(string(s))
	fmt.Println(string(vector))

}
