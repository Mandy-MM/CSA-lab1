package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
	fmt.Println(slice)
}

func mapSlice(f func(a int) int, slice []int) []int {
	var slice2 []int
	for _, num := range slice {
		slice2 = append(slice2, f(num))
	}
	return slice2
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	var array2 [5]int
	for i, num := range array {
		array2[i] = f(num)
	}
	return array2

}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	newSlice := intsSlice[1:3]
	intsArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(mapSlice(addOne, intsSlice))
	fmt.Println(mapSlice(square, newSlice))
	fmt.Println(mapArray(addOne, intsArray))
	double(intsSlice)
}
