package main

import "fmt"

func chain(left, right chan int) {
	left <- 1 + <-right
}

func longChain(size int) {

	leftMost := make(chan int)
	left := leftMost
	right := leftMost

	for i := 0; i < size; i++ {
		right = make(chan int)
		go chain(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftMost)
}
