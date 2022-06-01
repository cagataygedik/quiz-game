package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name:")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You can play the game")
	} else {
		fmt.Println("You can't play the game")
		return

	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better, RTX3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have?")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("Your scored %v out of %v ", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("Your scored: %v%%", percent)
}
