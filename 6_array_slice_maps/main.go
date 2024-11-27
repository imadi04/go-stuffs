package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Arrays-------")
	arr := [5]int{1, 2, 3, 4, 5}
	arr[0] = 4
	arr[4] = 8
	size := len(arr)
	fmt.Printf("Size is : %v\n", size)
	for i, v := range arr {
		fmt.Printf("index: %v , value: %v \n", i, v)
	}

	fmt.Println("\n-----Slices-------")
	my_slice := make([]string, 2) //declare a string slice of size 2

	my_slice[0] = "sdhf"
	my_slice[1] = "jsdh"
	my_slice = append(my_slice, "hello", "hi", "there")

	for i, v := range my_slice {
		fmt.Printf("index: %v , value: %v \n", i, v)
	}

	fmt.Println("\n-----Map-------")
	m := make(map[string]int)
	m["One"] = 1
	m["Two"] = 2
	m["Three"] = 3
	delete(m, "Two")
	for k, v := range m {
		fmt.Printf("key: %v , value: %v \n", k, v)
	}

}
