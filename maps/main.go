package main

import "fmt"

func main() {
	// map[key]valor
	ml := make(map[int]string)
	ml[1] = "a"
	ml[2] = "b"
	ml[3] = "c"

	fmt.Println(ml)
	fmt.Println(ml[1])

	delete(ml, 2)
	fmt.Println(ml)

	ml[1] = "A2"
	fmt.Println(ml)

	ml[7] = ""
	fmt.Println(ml[7])
	fmt.Println(ml[99]) //no existe y trae valor por defecto de string que es ""

	// se puede acceder y obtener un segundo valor que nos dice si existe o no

	valor, existe := ml[7]
	fmt.Println(valor, existe)
	valor2, existe2 := ml[9]
	fmt.Println(valor2, existe2)

	m2 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	fmt.Println(m2)
}
