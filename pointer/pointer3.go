package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Olga",
	}
}

func main() {
	person := &Person{
		Name: "Eugene",
	}
	fmt.Println(person.Name) // что отобразится после вызова? // Eugene
	changeName(person)
	fmt.Println(person.Name) // что отобразится после вызова? // Eugene
}
