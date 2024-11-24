package main

import "fmt"

type Person1 struct {
	Name string
}

func changeName1(person *Person1) {
	person = &Person1{
		Name: "Olga",
	}
}

func main() {
	person := &Person1{
		Name: "Eugene",
	}
	fmt.Println(person.Name) // что отобразится после вызова? // Eugene
	changeName1(person)
	fmt.Println(person.Name) // что отобразится после вызова? // Eugene
}
