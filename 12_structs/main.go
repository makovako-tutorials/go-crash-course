package main

import ("fmt"
		"strconv"
)

// define struct

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// method value receiver
func (p Person) greet()string {
	return "Hello, " + p.firstName + " " + p.lastName + " age " + strconv.Itoa(p.age) 
}

//method pointer receiver - we will change sth

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMerried(lastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = lastName

	}
}



func main() {
	//init person struct

	person1 := Person {
		firstName: "hello",
		lastName: "world",
		city:"New york",
		gender:"male",
		age:2,
	}

	person2 := Person {"hello", "world", "las vegas", "female", 4 }
	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println(person2.greet())
	person2.getMerried("worldovs")
	fmt.Println(person2.greet())
	


}