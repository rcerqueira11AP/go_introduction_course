package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	sum = 0
	for {
		if sum > 10000 {
			break
		}
		sum++
	}

	fmt.Println(sum)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range arr {
		fmt.Println("index: ", i, "- value: ", arr[i])
	}

	fmt.Println()
	for i, v := range arr {
		fmt.Println("index: ", i, "- value: ", v)
	}

	fmt.Println()

	//recorrer map
	map2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range map2 {
		fmt.Println("key: ", key, "- value: ", value)
	}

	fmt.Println()
	map3 := map[string][]int{
		"one":   nil,
		"two":   {4, 5, 6},
		"three": {7, 8, 9},
	}

	for key, value := range map3 {
		fmt.Println("key: ", key)
		for _, v := range value {
			fmt.Println("value: ", v)
		}
		fmt.Println()
	}

	// tarea
	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad
	for i, v := range array {
		array[i] = v + 20
	}

	fmt.Println("Los valores del array son: ", array)

	// var nums []int

	// for {
	// 	var input int

	// 	fmt.Print("Ingrese un numero: ")
	// 	_, err := fmt.Scanln(&input)
	// 	if err != nil {
	// 		fmt.Println("Error: ", err)
	// 		continue
	// 	}

	// 	if input == 0 {
	// 		break
	// 	}

	// 	nums = append(nums, input)

	// }

	// fmt.Println("Los numeros ingresados son: ", nums)

	code_list := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	var items []string

	for {
		var code string

		fmt.Print("Ingrese un codigo: ")
		_, err := fmt.Scanln(&code)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if code == "0" {
			break
		}

		value, found := code_list[code]
		if found {
			items = append(items, value)
		} else {
			items = append(items, "No encontrado")
		}

	}

	fmt.Println("Los items registrados son: ", items)
}
