package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less then %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is greater then %d\n", x, y)
	} else {
		fmt.Println("they are equal")
	}

	color := "red"

	switch color {
	case "red":
		fmt.Println("red")
	case "blue":
		fmt.Println("blue")
	default:
		fmt.Println("not blue or red")
	}
}