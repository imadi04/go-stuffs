package main

import (
	"fmt"
	"my_func/closures"
)

func greet(name string) string {
	return "Hello " + name
}

// A function that accepts another function as an argument
func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(x, y int) int {
	return x + y
}

func subtract(a, b int) int {
	return a - b
}

// A function that returns another function
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

// A variadic function that sums numbers
func sum_v(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// A function with regular and variadic parameters
func greet_v(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

func main() {

	fmt.Println("---------Assigning function to a variable---------")

	var myfun func(string) string // defining type of variable , myfun is of type 'function' with arg and return type as string

	myfun = greet // Assigning function to a variable
	fmt.Println(myfun("Aditya"))

	fmt.Println("----------passing Functions as Arguments----------")
	result := operate(3, 4, add)
	fmt.Println("Result:", result) // Output: Result: 7

	fmt.Println("-----------Returning Functions from a Function--------------")
	double := multiplier(2) // Create a function to double numbers
	triple := multiplier(3) // Create a function to triple numbers

	fmt.Println(double(5)) // Output: 10
	fmt.Println(triple(5)) // Output: 15

	fmt.Println("-----------Storing Functions in a Map or Slice-------------")
	// Store functions in a map
	operations := map[string]func(int, int) int{
		"add":      add,
		"subtract": subtract,
	}

	fmt.Println("Addition:", operations["add"](10, 5))         // Output: Addition: 15
	fmt.Println("Subtraction:", operations["subtract"](10, 5)) // Output: Subtraction: 5

	fmt.Println("----------variadic function------------")

	fmt.Println(sum_v(1, 2, 3, 4)) // Output: 10
	fmt.Println(sum_v(5, 10))      // Output: 15
	fmt.Println(sum_v())           // Output: 0 (no arguments passed)
	//// A function with regular and variadic parameters
	greet_v("Hello", "Alice", "Bob", "Charlie")

	fmt.Println("-----Basic example of closure---------")
	closures.My_closure_func()

	// closure with state
	counter1 := closures.Counter()
	counter2 := closures.Counter()

	fmt.Println(counter1()) // Output: 1
	fmt.Println(counter1()) // Output: 2
	fmt.Println(counter2()) // Output: 1
	fmt.Println(counter2()) // Output: 2

	//closure as function fsctory
	double_closure := closures.Multiplier(2) // Create a function to double a number
	triple_closure := closures.Multiplier(3) // Create a function to triple a number

	fmt.Println(double_closure(5)) // Output: 10
	fmt.Println(triple_closure(5)) // Output: 15
}
