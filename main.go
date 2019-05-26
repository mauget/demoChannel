package main

import (
	"fmt"
	"time"
)

// prints to console directly at intervals -- it produces side-effects -- no return value
func countDown() {

	for delay := 5; delay > 0; delay-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d ...\n", delay)
	}

	fmt.Println("(msg from countdown; I'm finishing)\n")
}

// returns a value to a channel after six seconds --  impure action is returning via an async channel instead of directly.
func echo(arg string, c chan string) {
	time.Sleep(6 * time.Second)
	c <- fmt.Sprintf("%s\n", arg)
}

func main() {
	go countDown()
	fmt.Println("(msg from main: the goroutine countdown is probably executing)\n")

	c := make(chan string)
	go echo("Hello, World, via six-second-delayed goroutine channel!", c)
	fmt.Print(<-c)
}
