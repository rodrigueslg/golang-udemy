package main

import (
	"fmt"
)

func main() {
	// ARRAY
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 100
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{11, 12, 13, 14, 15}
	fmt.Println(array3)

	// SLICE
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 16)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 1000
	fmt.Println(slice2)

	// ARRAY INTERNO
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
