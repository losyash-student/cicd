package main

import "fmt"

func minus(a, b int) int {
	return a - b
}

func summ(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(summ(1, 2))
	fmt.Println(minus(1, 2))
}
