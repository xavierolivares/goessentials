// Example of "if" statement
package main

import (
	"fmt"
)

func main() {
	x := 14

	if x > 100 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is not so big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a more than half of b")
	}
}
