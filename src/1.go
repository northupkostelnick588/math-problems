package main

import "math/rand"

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println(x + y)
}
