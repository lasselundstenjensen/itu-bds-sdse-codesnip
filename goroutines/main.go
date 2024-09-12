package main

import (
	"fmt"
	"time"
)

func main() { // main goroutine
	go PrintMessage("Hi from another function") // another goroutine
	go PrintMessage("Yay!")                     // another goroutine
	PrintMessage("Hello from this function")    // regular function call
	fmt.Println("Bye!")
}

func PrintMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
}
