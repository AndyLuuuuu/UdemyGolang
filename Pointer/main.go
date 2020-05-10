package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	jim := person{firstName: "jim", lastName: "anderson"}

	jimPointer := &jim
	jimPointer.updateName("Alex")

	jim.print()

	andy := person{firstName: "Andy", lastName: "Lu"}

	andy.updateName("Andrew")
	andy.print()
}
