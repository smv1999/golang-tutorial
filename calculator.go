package main

import "fmt"

func main() {
	var num1, num2, choice int
	fmt.Print("Enter the first number:")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number:")
	fmt.Scanln(&num2)

	fmt.Println("1. Addition 2. Subtraction 3. Multiplication 4. Division 5. Modulus")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println(num1 + num2)
	} else if choice == 2 {
		fmt.Println(num1 - num2)
	} else if choice == 3 {
		fmt.Println(num1 * num2)
	} else if choice == 4 {
		fmt.Println(num1 / num2)
	} else if choice == 5 {
		fmt.Println(num1 % num2)
	} else {
		fmt.Println("Invalid Choice!")
	}

}
