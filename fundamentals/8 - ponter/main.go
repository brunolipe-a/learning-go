package main

import "fmt"

func main() {
	value := 1

	pointer := &value

	*pointer++
	value++

	fmt.Println(pointer, *pointer, value)
}
