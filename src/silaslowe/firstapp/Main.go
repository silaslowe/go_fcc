package main

import "fmt"

// Cannot use := at package level. must use var i int = 42

func main() {
	var g int
	g = 55
	var h int = 56
	i := 42
	fmt.Println(i)
	fmt.Println(g)
	fmt.Println(h)
}