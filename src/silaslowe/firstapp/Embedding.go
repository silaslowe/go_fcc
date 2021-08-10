package main

import (
	"fmt"
	"reflect"
)	

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

type Aminal struct {
	Name string `required max: "100"`
	Origin string
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Origin)

	t := reflect.TypeOf(Aminal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}