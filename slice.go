package main

import  (
	"fmt"
	"slices"
	//"strings"
)

func main() {
	// Slice Integer
	numbers := [] int{5, 2, 9, 8, 1, 3}
	fmt.Println("Slice awal: ", numbers)

	//Menambahkan element
	numbers = append(numbers, 10)
	fmt.Println("Setelah ditambahkan : ", numbers)

	//Mengurutkan element
	sortedNumbers := slices.Clone(numbers)
	slices.Sort(sortedNumbers)
	fmt.Println("Setelah diurutkan: ", sortedNumbers)

	//Mencari element dalam slice
	target := 11
	index := slices.Index(numbers, target)
	if index != -1 {
		fmt.Printf("Element %d ditemukan pada index: %d\n", target, index)
	} else {
		fmt.Printf("Element %d tidak ditemukan\n", target)
	}
}