package main

func main() {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	variadicString(s...)
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
}
