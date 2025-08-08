package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5) // ring dengan 5 elemen

	// Mengisi data pada setiap node dalam ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next() // pindah ke node berikutnya
	}

	// Menampilkan seluruh data
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
