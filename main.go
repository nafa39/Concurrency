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

// function to print the alphabets from a to j
func printLetters1() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c\n", i)
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
func printLetters2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c\n", i)
	}
}

// produce and consume
func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("Produced:", i, "Buffered channel status:", len(ch))
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consume(ch chan int) {
	for i := range ch {
		fmt.Println("Consumed:", i, "Buffered channel status:", len(ch))
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	//task1
	// concurrency (goroutine)

	// calling the function printNumbers() in a goroutine
	go printNumbers1()

	// calling the function printLetters() in a goroutine
	go printLetters1()

	// sleep
	time.Sleep(300 * time.Millisecond)
	fmt.Println("=====================================")

	//task2
	// wait for the goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()

	// sleep
	time.Sleep(300 * time.Millisecond)
	fmt.Println("=====================================")

	//task3
	// channels
	ch := make(chan int)
	go produce(ch)
	consume(ch)

	// sleep
	time.Sleep(300 * time.Millisecond)
	fmt.Println("=====================================")

	//task4
	// buffered channels
	ch1 := make(chan int, 5)
	go produce(ch1)
	consume(ch1)

	// difference between task3 and task4
	// In task3, the channel is unbuffered, so the producer will block until the consumer reads the data from the channel.
	// In task4, the channel is buffered, so the producer will not block until the buffer is full.

	// sleep
	time.Sleep(300 * time.Millisecond)
	fmt.Println("=====================================")

}
