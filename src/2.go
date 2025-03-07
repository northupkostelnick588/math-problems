package main

import "math/rand"

func generateRandomMathProblem(n int) (int, int) {
	a := rand.Intn(n)
	b := rand.Intn(n)
	return a, b
}
