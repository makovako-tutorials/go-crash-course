package main

import "fmt"

func main() {
	// string, float43/64, int, bool ...
	// Using var
	var name string = "Brad"
	var age int = 32

	//Shorthand
	name2 := "lala"
	a, b := 1, 2
	fmt.Println(name, age, name2, a, b)
	fmt.Printf("%T\n", age)
}
