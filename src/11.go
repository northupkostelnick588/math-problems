package main

import "math"

func main() {
	x := 10
	y := 20
	z := math.Sqrt(x*x + y*y)

	fmt.Println("The length of the vector is", z)
}
