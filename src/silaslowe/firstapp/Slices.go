package main

import "fmt"

func main() {
	// a := []int{1,2,3}
	// b := a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// s := []int{1,2,3,4,5,6,7,8,9,10}
	// t := s[:] // slice of all elements
	// u := s[3:] // slice from 4th elemenet to end
	// v := s[:6] // slice first 6 elements
	// w := s[3:6] // slice the 4th, 5th and 6th elements

	// s[5] = 42
	// fmt.Println(s)
	// fmt.Println(t)
	// fmt.Println(u)
	// fmt.Println(v)
	// fmt.Println(w)

	// make function

	// a := make([]int, 3, 100)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	a := ([]int{})
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 2,3,4,5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, []int{21,31,41,51}...)
	// ... works like the spread operator in JS so we can pass an array into an array
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	// Removes first element of slice
	b := a[1:]
	fmt.Println(b)
	// Removes Element from end of slice
	c := a[:len(a) -1]
	fmt.Println(c)
	// Removes elements from inside the slice
	d := append(a[:2], a[4:]...)
	fmt.Println(d)
	fmt.Println(a)

}