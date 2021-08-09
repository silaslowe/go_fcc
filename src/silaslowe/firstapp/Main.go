// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// // Cannot use := at package level. must use var i int = 42

// // var block at package level
// var (
// 	actorName string = "Silas Lowe"
// 	companionName string = "Boba Fet"
// 	doctorNumber int = 3
// 	seasonNumber int = 11
// )

// func main() {
	
// 	// Three ways to decalare vars
// 	var g int
// 	g = 55
// 	var h int = 56
// 	i := 42


// 	// Type conversion between numeric values

// 	var s int = 99
// 	fmt.Printf("%v, %T\n", s, s)

// 	var l float32
// 	l = float32(s)
// 	fmt.Printf("%v, %T\n", l, l)
	

// 	fmt.Println(companionName)
// 	fmt.Println(doctorNumber)
// 	fmt.Println(h,i,g)

// 	// Type conversion from numeric to string

// 	var x int = 99
// 	fmt.Printf("%v, %T\n", x, x)

// 	var z string
// 	z = strconv.Itoa(s)
// 	fmt.Printf("%v, %T\n", z, z)

// }