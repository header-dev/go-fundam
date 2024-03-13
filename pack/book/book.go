package book

type Book struct {
	Name string
}

func New() Book {
	return Book{
		Name: "The Go Programming Language",
	}
}
