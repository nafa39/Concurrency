package main

import (
	"fmt"
	"sync"
	"time"
)

// function to print the numbers from 1 to 10
func printNumbers1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// function to print the numbers from 1 to 10
func printNumbers2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// function to print the alphabets from a to j
func printLetters1() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c\n", i)
	}
}

// function to print the alphabets from a to j
func printLetters2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c\n", i)
	}
}

// function sleep
func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	//task1
	// concurrency (goroutine)

	// calling the function printNumbers() in a goroutine
	go printNumbers1()

	// calling the function printLetters() in a goroutine
	go printLetters1()

	// sleep for 1 second
	sleep(1)
	fmt.Println("=====================================")

	//task2
	// wait for the goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()

}
