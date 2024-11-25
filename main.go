package main

import (
	"fmt"
)

// function to print the numbers from 1 to 10
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// function to print the alphabets from a to j
func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Println(string(i))
	}
}

func main() {
	//task1
	// concurrency (goroutine)

	// calling the function printNumbers() in a goroutine
	go printNumbers()

	// calling the function printLetters() in a goroutine
	go printLetters()

}
