package main

import "fmt"

type rectangle struct {
	w, l float64
}

func (r rectangle) Area() float64 {
	return r.w * r.l
}

type Book struct {
	Name   string
	Author string
}

func (b *Book) SetName(name string) {
	b.Name = name
}

func (b Book) String() string {
	return b.Name + " by " + b.Author
}

func main() {
	r := rectangle{w: 3, l: 4}

	fmt.Println(r.Area())

	b := Book{Name: "Harry Potter", Author: "J.K. Rowling"}
	fmt.Println(b.String())

	b.SetName("Wolverine")

	fmt.Println(b.String())

}
