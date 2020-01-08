package main

import "fmt"

func main() {
	// map - key value pair

	emails := make(map[string]string)

	emails2 := map[string]string {"bob":"bob@email.com"}

	//Assign key values

	emails["Bob"] = "bob@gmail.com"
	emails["Lolo"] = "lolo@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails2)



	//delete from map

	delete(emails, "Lolo")
	fmt.Println(emails)

}