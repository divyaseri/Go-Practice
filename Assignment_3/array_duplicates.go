package main

import "fmt"

func main() {

	var array = []int{4, 12, 9, 5, 7, 12, 4, 2, 7, 3, 5, 3, 9, 2}
	var arr []int

	fmt.Println("Array: ", array)
	var i int
	var j int

	for i = 0; i < len(array); i++ {
		for j = i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				arr = append(arr, array[i])
			}

		}
	}
	fmt.Println("Array with no duplicates: ", arr)
}
