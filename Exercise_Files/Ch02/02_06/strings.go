// Go strings
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// unit8 is a byte

	// Strings in Go are immutable
	// book[0] = 116;

	// Slice (start, end), 0 base, 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was 1/2 price!")

	// Multi line
	poem := `
	...
	Twenty years from now you will be more disappointed in the things you didnt do rather than the things you did do
	So throw off the bowlings, sail away from the safe harbor, catch the trade winds in your sails
	Explore. Dream. Discover.
	...
	`
	fmt.Println(poem)

}
