package main

import "fmt"

func main() {
	num := 5

	fmt.Println("Original value: ", num)
	modifyPointer(&num)

	fmt.Print(num)
}

func modifyPointer(num *int) {
	*num = 10
}
