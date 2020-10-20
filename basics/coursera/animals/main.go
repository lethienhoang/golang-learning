package main

import (
	"fmt"
)

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type Animal struct {
	Food, Locomotion, Sound string
}

func (a Animal) Eat() {
	fmt.Println(a.Food)
}

func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.Sound)
}

func change() *int {
	a := 10
	return &a
}

func main() {
	// data := map[string]Animal{
	// 	"cow":   Animal{"grass", "walk", "moo"},
	// 	"bird":  Animal{"worms", "fly", "peep"},
	// 	"snake": Animal{"mice", "slither", "hsss"},
	// }

	// var animalInterface AnimalInterface
	// for {
	// 	var command, name, requestType string
	// 	fmt.Print(">")
	// 	fmt.Scanln(&command)
	// 	fmt.Scanln(&name)
	// 	fmt.Scanln(&requestType)

	// 	if command == "newanimal" {
	// 		data[name] = data[requestType]
	// 		fmt.Println("Created it!")
	// 	} else if command == "query" {
	// 		animalInterface = data[name]
	// 		switch requestType {
	// 		case "eat":
	// 			animalInterface.Eat()
	// 			return
	// 		case "move":
	// 			animalInterface.Move()
	// 			return
	// 		case "speak":
	// 			animalInterface.Speak()
	// 			return
	// 		}
	// 	}
	// }
	var b *int
	fmt.Println("a (before) = ", b)
	b = change()
	fmt.Println("a (before) = ", b)
	*b = 11
	fmt.Println("a (before) = ", b)
}
