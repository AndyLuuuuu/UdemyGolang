package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	jim := person{firstName: "jim", lastName: "anderson"}
	jim.print()
}
