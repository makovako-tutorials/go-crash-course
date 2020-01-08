package main

import "fmt"

func main() {
	// Arays - fixed length and type

	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	fruitArr2 := [2]string {"Hello", "world"}
	fmt.Println(fruitArr2)

	fruitSLice := []string {"hello", "hello", "world"}
	fmt.Println(fruitSLice)
	fmt.Println(len(fruitSLice))
	fmt.Println(fruitSLice[1:3])
}