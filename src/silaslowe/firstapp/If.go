package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("Silas is true")
	}

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"CA": 323423,
		"TX": 234234,
		"LA": 6456456,
		"NY": 19675,
	}
	if pop, ok := statePopulations["TX"]; ok {
		fmt.Println(pop)
	}

	number := 70
	guess := 70
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You gt it!!")
	}


	//Dealing with floating point approximations
	
	myNum := 0.12353
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < .001 {
		fmt.Println("Samsies")
	} else {
		fmt.Println("Not samsies")
	}
}