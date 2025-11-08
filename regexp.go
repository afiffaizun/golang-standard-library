package main

import (
	"regexp"
	"fmt"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z	])`)

	fmt.Println(regex.MatchString("Apip"))
	fmt.Println(regex.MatchString("ap"))
	fmt.Println(regex.MatchString("apip"))


}