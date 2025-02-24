package main

import "fmt"

func main() {

	name := "Allen"
	age := 25
	pi := 3.14
	isGoFun := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("PI:", pi)
	fmt.Println("isGoFun", isGoFun)

	isGoFun = false

	fmt.Println("isGoFun", isGoFun)

	const sex string = "male"
	//常數不可修改
	//sex = female

}
