package main

import "fmt"

type operete func(int, int) int

func main() {
	fmt.Println(calc(getnumandoperete()))
}
func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	if b == 0 {
		fmt.Println("error")
		return 0
	}
	return a / b
}
func calc(a int, b int, op operete) int {
	return op(a, b)
}
func getnumandoperete() (int, int, operete) {
	var a, b int
	var oper string
	fmt.Print("input with order' num1 num2 operete(+ - * /):' ")
	fmt.Scan(&a, &b, &oper)
	switch oper {
	case "+":
		return a, b, add
	case "-":
		return a, b, sub
	case "*":
		return a, b, mul
	case "/":
		return a, b, div
	default:
		fmt.Println("error operete")
		return 0, 0, func(int, int) int { return 0 }
	}
}
