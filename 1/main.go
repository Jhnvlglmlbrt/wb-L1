package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func main() {
	person := Action{
		Human{
			name: "Андрей",
			age:  30},
	}
	fmt.Printf("New age: %d | new name: %s\n", person.age, person.name)
	person.changeName("Андрей (Антон)")
	person.changeAge(10)
	fmt.Printf("New age: %d | new name: %s\n", person.age, person.name)

	person.speak()
	person.Human.speak()
}

func (a *Action) changeAge(age int) {
	a.age = age
}

func (a *Action) changeName(name string) {
	a.name = name
}

func (a Action) speak() {
	fmt.Println("Action is speking")
}

func (h Human) speak() {
	fmt.Println("Human is speking")
}
