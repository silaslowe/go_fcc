// package main

// import (
// 	"fmt"
// )

// func main() {

// // & shows the address in memory 
// // * in the print statment it gives the value

// 	// var a int = 42
// 	// var b *int = &a
// 	// fmt.Println(a, *b)

// 	a := [3]int{1,2,3}
// 	b := &a[0]
// 	c := &a[1]
// 	fmt.Printf("%v %p %p\n", a,b,c)


// 	// Uninitialized pointes have a default value of nil

// 	var ms *myStruct
// 	ms = new(myStruct)
// 	(*ms).foo = 42
// 	fmt.Println(ms.foo)

// 	s := []int{1,2,3}
// 	t := s
// 	fmt.Println(s, t)

// 	s[1] = 42
// 	fmt.Println(s, t)


// }

// type myStruct struct {
// 	foo int
// }