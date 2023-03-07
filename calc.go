package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Привіт! Це калькулятор.")

	for {
		fmt.Print("Введіть вираз: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Помилка: Невірний формат введення")
			continue
		}

		input = strings.TrimSpace(input)

		values := strings.Split(input, " ")
		if len(values) != 3 {
			fmt.Println("Помилка: Невірний формат виразу")
			continue
		}

		num1, err := strconv.ParseFloat(values[0], 64)
		num2, err := strconv.ParseFloat(values[2], 64)

		if err != nil {
			fmt.Println("Помилка: Невірний формат чисел")
			continue
		}

		operate(values[1], num1, num2)
	}
}

func operate(operator string, num1, num2 float64) {
	var result float64
	var err error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result, err = divide(num1, num2)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Помилка: Невірна операція")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("помилка: Ділення на нуль неможливе")
	}
	return num1 / num2, nil
}
