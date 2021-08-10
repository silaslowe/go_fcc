// package main

// import "fmt"

// func main() {
// 	grades := [...]int{97,85,93}
// 	fmt.Printf("Grades %v", grades)

// 	var students [3]string
// 	fmt.Printf("Students %v\n", students)
// 	students[0] = "Lisa"
// 	students[2] = "Ahmed"
// 	students[1] = "Frank"
// 	fmt.Printf("Students %v\n", students)
// 	fmt.Printf("Student #2: %v\n", students[1])
// 	fmt.Printf("Students: %v, %v, %v\n", students[0], students[1], students[2])
// 	fmt.Printf("Students: %v\n", len(students))

// 	// b is a copy of a and not a pointer to the vlue in a
// 	// a := [...]int{1,2,3}
// 	// b := a
// 	// b[1] = 5
// 	// fmt.Println(a)
// 	// fmt.Println(b)

// 	// & makes be a pointer to the values in a and they are the same

// 	a := [...]int{1,2,3}
// 	b := &a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }