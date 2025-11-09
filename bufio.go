package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	//Membaca dengan Buffer
	text := "Ini contoh penggunaan bufio"
	reader := strings.NewReader(text)
	bufferedReader := bufio.NewReader(reader)

	//Membaca 20 byte pertama
	buffer := make([]byte, 20)
	n, _ := bufferedReader.Read(buffer)
	fmt.Printf("Membaca %d byte: %s\n", n, string(buffer[:n]))
	fmt.Println()

	//Membaca sisa data
	remaining, _ := bufferedReader.ReadString('\n')
	fmt.Printf("Sisa data: %s\n", remaining)

}