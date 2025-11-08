package main

import (
	//"encoding/base64"
	"encoding/json"
	//"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	Name   string `json:"name" xml:"name"`
	Age    int    `json:"age" xml:"age"`
	Email  string `json:"email" xml:"email"`
	IsAdmin bool   `json:"is_admin" xml:"is_admin"`	
}

func main() {
	fmt.Println("Encoding Example")


	jsonExample()

	//xmlExample()

	//base64Example()
}

func jsonExample() {
	person := Person{
		Name:   "Apip",
		Age:	30,
		Email:  "apipgmail.com",
		IsAdmin: true,
	}

	// Encoding Struct -> JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON Encoded:", string(jsonData))

	jsonPretty, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nJson Pretty")
	fmt.Println(string(jsonPretty))
}

	