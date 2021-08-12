package main

import (
	"fmt"
	"log"
)

func main() {

	// Last in/first out with defer
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// res, err := http.Get("http://google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err!= nil {
	// 	log.Fatal()
	// }
	// fmt.Printf("%s", robots)

	// Defer evaluated the value of variable at the time it was called: result is a = "state"

	// a := "start"
	// defer fmt.Println(a)
	// a = "end"
	
	// Panic runs after defered 

	fmt.Println("start")
	defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
			}
	}()
	panic("something bad happened")
	fmt.Println("end")

	// This will fail if another temrinal opens and tries to run this on the same port

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Hello Werks"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }
}