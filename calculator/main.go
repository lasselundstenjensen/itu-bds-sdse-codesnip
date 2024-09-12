package main

import "fmt"

func main() {
	var operation string
	var number1 float64
	var number2 float64

	fmt.Println("CALCULATOR, ITU GO COURSE V1.0")
	fmt.Println("==============================")

	fmt.Println("Choose an operation: add, subtract, multiply, divide")
	fmt.Scanf("%s", &operation)
	fmt.Println("What is the first number?")
	fmt.Scanf("%f", &number1)
	fmt.Println("What is the second number?")
	fmt.Scanf("%f", &number2)

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}
}
