package main

import (
	"fmt"
	"path"
	
)

func main() {
	//Menggabungkan path
	// dir := "home"
	// user := "smart"
	// file := "documents.txt"

	//Menggabungkan beberapa bagian path
	// fullPath := filepath.Join(dir, user, file)
	// fmt.Printf("Dir: %s User: %s, File: %s\n", dir, user, file)
	// fmt.Println("Path Gabungan: ", fullPath)

	// //Memisahkan dir dan file
	// filePath := "home/user/document.txt"
	// dirPath, fileName := path.Split(filePath)
	// fmt.Printf("Path lengkap: %s\n", filePath)
	// fmt.Printf("Dir: %s\n", dirPath)
	// fmt.Printf("File: %s\n", fileName)

	//Contoh dengan path root
	rootPath := path.Dir("/home/user")
	fmt.Printf("Path root: %s\n", rootPath)
}