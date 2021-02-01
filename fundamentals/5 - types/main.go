package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100
	fmt.Println("'a' is", a)
	fmt.Println("Type of 'a':", reflect.TypeOf(a))

	b := 11.12
	fmt.Println("\n'b' is", b)
	fmt.Println("Type of 'b':", reflect.TypeOf(b))

	c := float32(11.12)
	fmt.Println("\n'c' is", c)
	fmt.Println("Type of 'c':", reflect.TypeOf(c))

	// unsigned: uint8, uint16, uint32, uint64
	d := byte(255)
	fmt.Println("\n'd' is", d)
	fmt.Println("Type of 'd':", reflect.TypeOf(d))

	e := false
	fmt.Println("\n'e' is", e)
	fmt.Println("not 'e' is", !e)
	fmt.Println("Type of 'e':", reflect.TypeOf(e))

	f := "Hello, my name is Bruno"
	fmt.Println("\n'f' is", f)
	fmt.Println("lenght of 'f' is", len(f))
	fmt.Println("Type of 'f':", reflect.TypeOf(f))

	g := `Hello,
	my
	name
	is
	Bruno`

	fmt.Println("\n'g' is", g)
	fmt.Println("lenght of 'g' is", len(g))
	fmt.Println("Type of 'g':", reflect.TypeOf(g))
}
