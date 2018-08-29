package main

import (
	"fmt"
	"math/rand"
	"time"
)

func borningMain() {
	rand.Seed(time.Now().UnixNano())

	// c1 := fanIn(borning("joe"), borning("ann"))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c1)
	// }

	c2 := make(chan string)

	go timeOutChannel(2, c2)
	sleepTime := rand.Intn(5)
	fmt.Printf("I'm going to sleep for %d seconds \n", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c2 <- "I'm awake now"
	close(c2)
	time.Sleep(1 * time.Second)
}

func timeOutChannel(timeout int, c <-chan string) {
	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Println("Timeout. No message received.")
	}

	fmt.Println(<-c + " is received but discarded")
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
	// go func() {for { c <- <-c1}}()
	// go func() {for {c <- <-c2}}()

	go func() {
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
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
