package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2

	area := PI * math.Pow(radius, 2)

	fmt.Println("Area of the circle:", area)

	const (
		a = 2
		b = 3
	)

	var (
		c = 4
		d = 5
	)

	fmt.Println(a, b, c, d)

	var e, f = true, false

	fmt.Println(e, f)

	g, h, i := 2, true, "ops!"

	fmt.Println(g, h, i)
}
