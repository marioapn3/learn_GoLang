package main

import (
	"fmt"
)

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(len(fruits))
	//array multidimensi
	var numbers = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers", numbers)
	fmt.Println("numbers[0]", numbers[0])
	fmt.Println("numbers[0][1]", numbers[0][1])
	fmt.Println(len(numbers))

	//type slices
	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits2[0])
	fmt.Println(fruits2[1])
	fmt.Println(len(fruits2))
	//slice multidimensi	
	var numbers2 = [][]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers2", numbers2)
	fmt.Println("numbers2[0]", numbers2[0])
	fmt.Println("numbers2[0][1]", numbers2[0][1])
	fmt.Println(len(numbers2))
	
}
