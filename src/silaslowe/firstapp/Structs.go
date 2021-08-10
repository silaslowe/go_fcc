package main

import "fmt"

type Doctor struct {
	Number int
	ActorName string
	Companions []string
}

func main() {
	aDoctor := Doctor {
		Number: 3,
		ActorName: "Jon Pertwee",
		Companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.ActorName)
	fmt.Println(aDoctor.Companions[1])

	// Anonomous struct

	bDoctor := struct{name string}{name: "Silas Lowe"}
	anotherDoctor := bDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(bDoctor)
	fmt.Println(anotherDoctor)


}