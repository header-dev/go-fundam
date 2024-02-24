package main

import (
	"fmt"
	"strings"
)

func main() {

	m := map[string]string{}
	if m == nil {
		fmt.Println("m is nil")
	}

	m["a"] = "alpha"

	fmt.Println(m)

	s := "Apple Banana Apple Banana apple"
	fmt.Println(wordCount(s))

}

func wordCount(s string) map[string]int {

	split := strings.Split(s, " ")
	result := map[string]int{}
	for i := 0; i < len(split); i++ {
		result[split[i]] = result[split[i]] + 1
	}
	return result

}
