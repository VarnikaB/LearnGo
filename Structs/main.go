package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//var alex person -> create a structure alex with values as  "" ""
	//default values: int, float: 0;;;bool: false;;;string: ""
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v\n", alex) // prints the field names with values

	var alice person
	alice.firstName = "Alice" // = not := as the field are already declared
	alice.lastName = "Sanderson"
	fmt.Printf("%+v\n", alice)
}
