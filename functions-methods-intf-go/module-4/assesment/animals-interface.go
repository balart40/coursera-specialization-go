package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	sound      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (receivedAnimal Animal) Eat() {
	fmt.Println(receivedAnimal.food)
}

func (receivedAnimal Animal) Move() {
	fmt.Println(receivedAnimal.locomotion)
}

func (receivedAnimal Animal) Speak() {
	fmt.Println(receivedAnimal.sound)
}

func main() {
	var command, providedAnimal, infoReq string
	var animalI animalInterface

	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		fmt.Print(">")
		_, err := fmt.Scan(&command, &providedAnimal, &infoReq)

		if err != nil {
			fmt.Println("Error scanning vars input")
		}

		if command == "query" {
			animalI = animals[providedAnimal]
			if infoReq == "eat" {
				animalI.Eat()
			} else if infoReq == "move" {
				animalI.Move()
			} else if infoReq == "speak" {
				animalI.Speak()
			}
		} else if command == "newanimal" {
			animals[providedAnimal] = animals[infoReq]
			fmt.Println("Created it!")
		}
	}
}
