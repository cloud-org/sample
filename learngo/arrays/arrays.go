package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	//for i := range arr3 {
	//	fmt.Println(arr3[i])
	//}

	printArray(&arr1)
	//printArray(arr2)
	printArray(&arr3)

	fmt.Println(arr1, arr3)
}
