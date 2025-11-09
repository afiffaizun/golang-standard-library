package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	source:= strings.NewReader("Hello, World!")
	destination := &strings.Builder{}

	bytesWritten, err := io.Copy(destination, source)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Jumlah bytes yang disalin: %d\n", bytesWritten)
		fmt.Println("Data yang disalin: ", destination.String())
	}
}