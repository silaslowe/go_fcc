// package main

// import (
// 	"fmt"
// )

// func main() {
// 	switch 5 {
// 	case 1, 5, 10:
// 		fmt.Println("1, 5 or 10")
// 	case 2, 6:
// 		fmt.Println("2 or 6")
// 	default: 
// 	fmt.Println("needer")
// 	}

// 	switch i := 2 + 3;i {
// 	case 1, 5, 10:
// 		fmt.Println("1, 5 or 10")
// 	case 2, 6:
// 		fmt.Println("2 or 6")
// 	default: 
// 	fmt.Println("needer")
// 	}
// 	j := 10
// 	switch {
// 	case j <= 10:
// 		fmt.Println("less than or the same as 10")
// 		// fallthrough keyword avoids a break which is implied in GO
// 		fallthrough
// 	case j <= 20:
// 		fmt.Println("less than or the same as 20")
// 	default:
// 		fmt.Println("Bigger than 20")	
// }


// 	var i interface{} = "hi"
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is a int")
// 		// break
// 		// fmt.Println("this will pring too")
// 	case float64:
// 		fmt.Println("i is a 64")
// 	case string:
// 		fmt.Println("i is a string")
// 		//This must match the type of array passed as the tested value ie not [2]string
// 	case [3]int:
// 		fmt.Println("i is a [3]int")
// 	default:
// 		fmt.Println("i is a another type")
// 	}
// }