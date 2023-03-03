package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ведіть вираз")
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("Відповідь: %d\n", Add(txt))
		fmt.Println("Ведіть вираз")
	}
}

func Add(txt string) int {
	resArr := strings.Split(txt, "+")
	var expr1 string
	var expr2 string
	if len(resArr) == 2 {
		expr1 = resArr[0]
		expr2 = resArr[1]
	}
	res, err := AddExpression(expr1, expr2)
	if err != nil {
		err := fmt.Errorf("помилка")
		fmt.Println(err)
		return 0
	}
	return res
}

func AddExpression(expr1 string, expr2 string) (int, error) {
	//TODO Переписати так щоб можна було необмежену кількість пераметрів передавати з одним знаком.
	res1, err := strconv.Atoi(expr1)
	res2, err := strconv.Atoi(expr2)
	return res1 + res2, err
}
