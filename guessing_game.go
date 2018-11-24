package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var quit = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)
	fmt.Println("Secret number is:", secretNumber)

	for quit != true {
		fmt.Printf("Please enter a number: ")
		fmt.Scan(&userInput)
		if userInput == secretNumber {
			fmt.Println("You won!")
			quit = true
		} else if userInput < secretNumber {
			fmt.Println("Too Low...")
		} else if userInput > secretNumber {
			fmt.Println("Too High...")
		}
	}
}
