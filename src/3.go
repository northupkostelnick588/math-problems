package main

import "math/rand"

func GenerateRandomMathProblem() string {
    // Generate two random numbers between 1 and 10
    num1 := rand.Intn(10) + 1
    num2 := rand.Intn(10) + 1

    // Generate a random operation (+, -, *, /)
    op := []string{"+", "-", "*", "/"}[rand.Intn(4)]

    // Return the math problem in string form
    return num1 + " " + op + " " + num2 + " = ?"
}