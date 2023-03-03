package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	txt := scanner.Text()
	//	AddText(txt)
	//	fmt.Printf("Эхо: %s\n", txt)
	//	if txt == "0" {
	//		return
	//	}
	//}

	expression := "k+k"
	res, err := AddText(expression)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func AddText(txt string) (int, error) {
	resArr := strings.Split(txt, "+")
	res, err := Add(resArr[0], resArr[1])
	return res, err
}

func Add(expr1 string, expr2 string) (int, error) {
	res1, err := strconv.Atoi(expr1)
	res2, err := strconv.Atoi(expr2)
	return res1 + res2, err
}
