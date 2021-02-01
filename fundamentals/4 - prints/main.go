package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141592

	fmt.Printf("The value of x is %.2f.", x)

	a := 1
	b := 1.9999
	c := true
	d := "ops!"

	fmt.Printf("\n%d, %f, %.2f, %t, %s", a, b, b, c, d)
}
