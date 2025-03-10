package main

import (
	"encoding/json"
	"fmt"
)

type Person map[string]string

func main() {
	var nameStr string
	var addrStr string
	newInputPerson := make(Person)
	// get name
	fmt.Println("Provide a name:")
	_, err := fmt.Scan(&nameStr)
	if err != nil {
		fmt.Println("Error scanning name input")
	}
	newInputPerson["name"] = nameStr
	// get address
	fmt.Println("Provide a address:")
	_, err = fmt.Scan(&addrStr)
	if err != nil {
		fmt.Println("Error scanning address input")
	}
	newInputPerson["address"] = addrStr
	marshalJson(newInputPerson)
}

func marshalJson(person Person) {
	// Mashal the map to Json
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling map to JSON:", err)
		return
	}
	// print
	fmt.Println(string(jsonBytes))
}
