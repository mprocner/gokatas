// A channel allows for communication and synchronization between goroutines.
//
// Level: beginner
// Topics: goroutines, channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		c <- fmt.Sprintf("%s", "xD")
		c <- fmt.Sprintf("%s", "xs")
	}()
	go say("blah", c)
	for i := 0; i < 30; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("end")
}

func say(msg string, c chan string) {
	for i := 0; i < 25; i++ {
		c <- fmt.Sprintf("%s, %d", msg, i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
	close(c)
}
