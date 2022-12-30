package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
	fmt.Print((sub(10, 5)))
}

func Soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}
