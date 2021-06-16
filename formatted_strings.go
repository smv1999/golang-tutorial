package main

import "fmt"

func main() {
	age := 21
	name := "Vaidhyanathan S M"
	// %v - default format, %q - quotes, %T - type of
	fmt.Printf("my age is %d and my name is %s \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %.2f points \n", 225.55867)

	// Sprintf - save formatted strings
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is ", str)
}
