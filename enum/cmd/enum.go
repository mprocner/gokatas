// Enum shows how to use enumerated type.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/enum"
)

func main() {
	d := enum.South
	fmt.Print(d)
	switch d {
	case enum.North:
		fmt.Println(" goes up.")
	case enum.South:
		fmt.Println(" goes down.")
	case enum.NorthEast:
		fmt.Println(" goes up and right.")
	default:
		fmt.Println(" stays put.")
	}
}
