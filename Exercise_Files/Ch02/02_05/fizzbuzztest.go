// "for" loop examples
package main

import (
	"fmt"
)

func main() {

	fmt.Println("-----")

	for i := 1; i < 16; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue;
		} else if i%3 == 0 {
			fmt.Println("fizz")
			continue;
		} else if i%5 == 0 {
			fmt.Println("buzz")
			continue;
		} else {
			fmt.Println(i)
		}
	}
	
	fmt.Println("-----")

}
