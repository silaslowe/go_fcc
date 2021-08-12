package main

import "fmt"

func main() {
// 		for i:= 0; i < 5; i++ {
// 			sayMessage("Hello Go!", i)
// 		}
// }

// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is: ", idx)

// greeting := "Hello"
// name := "Stacu"
// sayGreeting(greeting, name)
// fmt.Println(name)
// }

// // Syntax for parameters that share data type
// // Syntax for passing by value-will not change data as a copy was made
// func sayGreeting(greeting, name string) {
// 	fmt.Println(greeting, name)

// // Syntax for passing by reference-does not make a copy and changes the data at the datas address in memory

// 	greeting := "Hello"
// 	name := "Stacu"
// 	sayGreeting(&greeting, &name)
// 	fmt.Println(name)
// 	}

// 	func sayGreeting(greeting, name *string) {
// 		fmt.Println(*greeting, *name)
// 		*name = "Ted"
// 		fmt.Println(*name)

// Variatic Parameter: There can be only one

	sum("The sum is: ", 10,2,3,4,5)
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
			result += v
	}
	fmt.Println(msg, result)
}