package main

import "fmt"

func main() {

// // Methods are functions with a specific context and that is given in the (g, greeter) params before the method is named
// // This is by value and so the method cannot change the value of the properties in the struct

// 	g := greeter{
// 		greeting: "Hello",
// 		name: "Go",
// 	}
// 	g.greet()

// }
// type greeter struct {
// 	greeting string
// 	name string
// }

// func (g greeter) greet() {
// 	fmt.Println(g.greeting, g.name)

	g := greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
	fmt.Println("This is the name:", g.name)

}
type greeter struct {
	greeting string
	name string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Heyo"
}