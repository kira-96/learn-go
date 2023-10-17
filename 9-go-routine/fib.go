package main

import (
	"fmt"
	"time"
)

func fib(ch chan<- int, q <-chan bool) {
	x, y := 1, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func printFib() {
	start := time.Now()

	command := ""
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		fib(ch, quit)
	}()

	for {
		fmt.Println(<-ch)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
