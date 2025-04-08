package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	sound      string
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
	var animal, infoReq string

	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		fmt.Print(">")
		_, err := fmt.Scan(&animal, &infoReq)

		if err != nil {
			fmt.Println("Error scanning vars input")
		}

		// Check for exit condition
		if animal == "x" || infoReq == "x" || (animal == "x" && infoReq == "x") {
			fmt.Println("Exiting...")
			break
		}
		if infoReq == "eat" {
			animals[animal].Eat()
		} else if infoReq == "move" {
			animals[animal].Move()
		} else if infoReq == "speak" {
			animals[animal].Speak()
		}
	}
}
