// Counter creates an integer counter, increments it, prints it and resets it.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/counter"
)

func main() {
	c := new(counter.Counter)
	c.Increment()
	fmt.Println(c.N())
	c.Reset()
	c.Increment()
	c.Increment()
	fmt.Println(c.N())

	s := new(counter.Summer)
	s.Add(1)
	s.Add(5)
	fmt.Println(s.Value())
	s.Reset()
	fmt.Println(s.Value())
}
