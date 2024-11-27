package closures

import (
	"fmt"
)

func My_closure_func() {
	message := "Hello"

	// Closure capturing the outer variable
	printMessage := func() {
		fmt.Println(message)
	}

	printMessage() // Output: Hello

	// Modifying the outer variable
	message = "Hi"
	printMessage() // Output: Hi
}

// closure with state
// A function returning a closure
func Counter() func() int {
	count := 0 // Captured variable

	return func() int {
		count++
		return count
	}
}

// closure as Function factory
func Multiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}
