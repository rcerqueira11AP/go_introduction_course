package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println(myIntVar)

	var myUintVar uint // solo puede ser positivo
	myUintVar = 12

	fmt.Println(myUintVar)

	var myStringVar string
	myStringVar = "Hello, World!"

	fmt.Println(myStringVar)
	fmt.Println("My variable address is: ", &myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println(myBoolVar)

	myIntVar2 := 12
	fmt.Println(myIntVar2)

	myStringVar2 := "Hello, World!"
	fmt.Println(myStringVar2)

	const myFirstConst = "a12"
	fmt.Println(myFirstConst)

	const myIntCont int = 12
	fmt.Println(myIntCont)

	fmt.Println()

	var my8BitInt int8
	fmt.Printf("Int default value: %d\n", my8BitInt)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitInt, unsafe.Sizeof(my8BitInt), unsafe.Sizeof(my8BitInt)*8)
	var my16BitInt int16
	fmt.Printf("Int default value: %d\n", my16BitInt)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitInt, unsafe.Sizeof(my16BitInt), unsafe.Sizeof(my16BitInt)*8)

	var my32BitInt int32
	fmt.Printf("Int default value: %d\n", my32BitInt)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitInt, unsafe.Sizeof(my32BitInt), unsafe.Sizeof(my32BitInt)*8)

	var my64BitInt int64
	fmt.Printf("Int default value: %d\n", my64BitInt)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitInt, unsafe.Sizeof(my64BitInt), unsafe.Sizeof(my64BitInt)*8)

	var myFloat32Var float32
	fmt.Printf("Float default value: %f\n", myFloat32Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)
	var myFloat64Var float64
	fmt.Printf("Float default value: %f\n", myFloat64Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	myStringVar5 := `Hello, World!
	How are you?
	I'm fine, thank you!`

	fmt.Println(myStringVar5)

	{
		// scope variable que solo estan dentro del bloque
		fmt.Println()
		floatVar := 12.0
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)

		floatVarString := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatVarString, floatVarString)

		//String a entero se usa package strconv
		intVal1, err := strconv.ParseInt("123123", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)
		//String a entero se usa package strconv
		intVal2, err := strconv.ParseInt("aaa123", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)

		floatVar1, _ := strconv.ParseFloat("123.123", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)
	}
}
