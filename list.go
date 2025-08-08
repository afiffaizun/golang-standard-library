package main

import (
	"fmt"
	"container/list"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Muhammad")
	data.PushBack("apip")
	data.PushBack("faizun")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e =e.Next() {
		fmt.Println(e.Value)
	}
}