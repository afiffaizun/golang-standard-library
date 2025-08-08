package main

import (
	"fmt"
	"math"
)

func main() {
	a := 16.2
	b := 9.4

	//Max
	fmt.Println("Nilai maksimum dari", a, "dan", b, "=", math.Max(a, b))

	//Min
	fmt.Println("Nilai minimum dari", a, "dan", b, "=", math.Min(a, b))

	//Pembulatan
	c := 3.5
	fmt.Println("Pembulatan ke bawah", c, "=",math.Floor(c))
	fmt.Println("Pembulatan ke atas", c, "=",math.Ceil(c))
	fmt.Println("Pembulatan ke terdekat", c, "=",math.Round(c))

}