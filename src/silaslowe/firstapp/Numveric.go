// package main

// import "fmt"

// func main() {
// 	// int is int 32
// 	// int8 -128 and 127 
// 	// int16 -32k 32k
// 	// int32 -2B 2B
// 	// int 64 -9Q 9Q

// 	n:= 42
// 	fmt.Printf("%v, %T\n", n, n)

// 	var o uint16 = 42
// 	fmt.Printf("%v, %T\n", o, o)

// 	// bit shifting
// 	a := 8  // 2^3
// 	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 or 64
// 	fmt.Println(a >> 3) // 2^3 * 2^3 = 2^0 or 1

// 	// Strings are arrays of bytes in GO. Here is how to access the value of a given index

// 	s := "This is a string"
// 	fmt.Printf("%v, %T\n", string(s[3]), s[3])

// 	//Concat
	
// 	x := "This is a string"
// 	y := " and so is this"
// 	fmt.Printf("%v, %T\n", x + y, x + y)

// 	// Returns array of ascii codes for each letter

// 	z := "This is a string"
// 	e := []byte(z)
// 	fmt.Printf("%v, %T\n", e, e)


// }