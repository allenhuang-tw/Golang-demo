package main

import "fmt"

func main() {
	sum := add(155, 686)

	fmt.Print(sum)
}

func add(a int, b int) int {
	return a + b
}
