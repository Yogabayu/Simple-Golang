package main

import (
	"fmt"
)

func main() {
	var operator string
	var number1, number2 float64

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&number1)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&number2)

	result := calculate(number1, number2, operator)
	fmt.Println("Hasil:", result)
}

func calculate(num1, num2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak valid")
			result = 0
		}
	default:
		fmt.Println("Error: Operator tidak valid")
		result = 0
	}

	return result
}
