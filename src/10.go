package main

import "math/rand"

func main() {
	// generate a random number between 1 and 10
	num := rand.Intn(10) + 1
	fmt.Println("Generated number:", num)
}
