// main.go
package main

import (
	"fmt"
	"src/sum/src/sum"
)

func main() {
	// Taking input num
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// Calling sum function
	result := sum.Add(num1, num2)
	fmt.Println("Sum:", result)
}
