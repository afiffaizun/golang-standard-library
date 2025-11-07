package main

import (
	"fmt"
	"time"
)

func main() {
	
	// Getting current time
	now := time.Now()

	fmt.Println("Current timt: %v", now)

	// Creating spesific time
	spesificTime := time.Date(2024, time.December, 25, 10, 30, 0, 0, time.UTC)
	fmt.Println("tes: ", spesificTime)

	//Time
	fmt.Println("Year: ", now.Year())
	fmt.Println("Mont: ", now.Month())
	fmt.Println("Day: ", now.Day()) 
	fmt.Println("Hour: ", now.Hour())

 

}