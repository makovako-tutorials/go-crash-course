package main

import "fmt"

func main() {
	ids := []int {1,2,3,45,22}

	for i, id := range ids {
		fmt.Printf("%d - id %d\n",i,id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)

	emails2 := map[string]string {"bob":"bob@email.com", "lolo":"lolo@email.com"}
	
	for k,v := range emails2 {
		fmt.Printf("%s, %s\n",k,v)
	}

}