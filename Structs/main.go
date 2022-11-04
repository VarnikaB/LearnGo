package main

import "fmt"

type contact struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	//var alex person -> create a structure alex with values as  "" ""
	//default values: int, float: 0;;;bool: false;;;string: ""
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contact{
			email:   "alex.anderson@gmail.com",
			zipcode: 564091,
		},
	}
	fmt.Println(alex)
	fmt.Printf("%+v\n", alex) // prints the field names with values

}
