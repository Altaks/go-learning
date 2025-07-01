package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a int, b int, c int) int {
	return a + b + c
}

func main() {
	res := plus(2, 3)
	fmt.Println("2 + 3 =", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
