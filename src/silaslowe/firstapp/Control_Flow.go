package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Last in/first out with defer
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err!= nil {
		log.Fatal()
	}
	fmt.Printf("%s", robots)

	// Defer evaluated the value of variable at the time it was called: result is a = "state"

	a := "start"
	defer fmt.Println(a)
	a = "end"
	

}