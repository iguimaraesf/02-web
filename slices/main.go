package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	// slice = append(slice, 10, 20, 30)
	// fmt.Println(slice)
	// fmt.Println(cap(slice))
	// slice = append(slice, -9)
	// fmt.Println(slice)
	// fmt.Println(cap(slice))

	// for i := 0; i < 20; i++ {
	// 	slice = append(slice, i)
	// 	fmt.Println("Len: ", len(slice), " Cap: ", cap(slice))
	// }
	sliceTeste := slice
	slice[0] = 10
	slice = append(slice, 7, 2, 99, 1024)
	slice[1] = -10
	slice[0] = -100
	fmt.Println(sliceTeste)
	fmt.Println(slice)

	sliceStr := []string{
		"oi",
		"mundo",
		"você",
		"está",
		"bem?",
	}
	fmt.Println(sliceStr)
	fmt.Println(sliceStr[:2])  // oi mundo
	fmt.Println(sliceStr[1:3]) // mundo você
	fmt.Println(sliceStr[2:])  // você está bem?
}
