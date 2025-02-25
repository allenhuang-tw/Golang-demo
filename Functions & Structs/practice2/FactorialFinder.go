package main

import "fmt"

func main() {
	factorial := findFactorial(5)

	fmt.Print(factorial)
}

func findFactorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * findFactorial(n-1)
}
