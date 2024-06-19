package main

import (
	"fmt"
)

// generateOddNumbers sends odd numbers from 1 to 19 to the provided channel
// and closes the channel when done.
func generateOddNumbers(oddChan chan<- int) {
	for i := 1; i <= 20; i += 2 {
		oddChan <- i
	}
	close(oddChan)
}

// generateEvenNumbers sends even numbers from 2 to 20 to the provided channel
// and closes the channel when done.
func generateEvenNumbers(evenChan chan<- int) {
	for i := 2; i <= 20; i += 2 {
		evenChan <- i
	}
	close(evenChan)
}

func main() {
	// Create channels for odd and even numbers
	oddChan := make(chan int)
	evenChan := make(chan int)

	// Start goroutines to generate odd and even numbers
	go generateOddNumbers(oddChan)
	go generateEvenNumbers(evenChan)

	// Loop from 1 to 20 and print numbers in sequential order
	// Alternating between odd and even channels
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			// Receive and print an odd number
			fmt.Println(<-oddChan)
		} else {
			// Receive and print an even number
			fmt.Println(<-evenChan)
		}
	}
}
