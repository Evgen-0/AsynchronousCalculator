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
		fmt.Printf("Відповідь: %d\n", NoName(txt))
		fmt.Println("Ведіть вираз")
	}
}

func NoName(txt string) int {
	var expr1 string
	var expr2 string
	var res int
	resArrPlus := strings.Split(txt, "+")
	resArrMinus := strings.Split(txt, "-")
	if len(resArrPlus) == 2 {
		expr1 = resArrPlus[0]
		expr2 = resArrPlus[1]
		res = Add(expr1, expr2)
	}
	if len(resArrMinus) == 2 {
		expr1 = resArrMinus[0]
		expr2 = resArrMinus[1]
		res = Minus(expr1, expr2)
	}
	return res
}

func Minus(expr1, expr2 string) int {
	res1, res2 := ToInt(expr1, expr2)
	return res1 - res2
}

func Add(expr1 string, expr2 string) int {
	res1, res2 := ToInt(expr1, expr2)
	return res1 + res2
}

func ToInt(expr1, expr2 string) (int, int) {
	var res1 int
	var res2 int
	var err error
	res1, err = strconv.Atoi(expr1)
	res2, err = strconv.Atoi(expr2)
	if err != nil {
		fmt.Println(fmt.Errorf("Помилка"))
	}
	return res1, res2
}
