package main

import (
	"fmt"
)

const(
	// iota increments and states at 0. _ says to not hold the 0 in memory
	_ = iota +5
	cat
	dog
	snake

)

func main() {
	// lowercase start to const means it is scoped to this package and a capital letter to start meas it will be exported globally
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	fmt.Printf("%v\n", cat)

}