package main

import (
	"fmt"

	"github.com/header-dev/demo/book"
)

func main() {
	b := book.New()

	fmt.Printf("%T %v\n", b, b)

	b.Name = "The Go Programming Language 2nd"
	fmt.Printf("%T %v\n", b, b)

}
