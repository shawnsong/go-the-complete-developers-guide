package main

import (
	"fmt"
	"math/rand"
	"time"
)

func borningMain() {
	c := fanIn(borning("joe"), borning("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

/**
 * Channels as a handle on a service
 * The problem with this function is that ann is blocked until joe has some value
 * to print. If ann is more talktive than joe then ann can't speak until joe speaks
 */
func borning1() {
	// joe is a service
	joe := borning("joe is borning")
	// ann is a service
	ann := borning("ann is borning")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You are borning. I'm leaving")

}

func fanIn(c1 <-chan string, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func borning(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- msg
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
