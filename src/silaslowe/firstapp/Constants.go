package main

import (
	"fmt"
)

func main() {
	// lowercase start to const means it is scoped to this package and a capital letter to start meas it will be exported globally
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
}