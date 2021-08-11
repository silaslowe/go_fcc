package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for s, t := 0, 0; s < 5; s, t = s+1, t+2 {
		fmt.Println(s, t)
	}


	// While loop for GO
	x := 0
	for x < 11 {
		fmt.Println(x)
		x++
	}

	w := 0
	for {
		fmt.Println(w*7)
		w++
		if w == 5 {
			break
		} 
	}

	// Nested for loop
Loop: 
	for m := 1; m <= 4; m++ {
		for n := 1; n <=3; n++ {
			fmt.Println(m * n)
			if m * n >= 3 {
				break Loop
			}
		}
	}


	// Range keyword makes key value pair output
	b := []string{"Dog","Cat","Mouse"}
	for k, v := range b {
		fmt.Println(k, v)
	}

	c := "Silas Lowe"
	for k, v := range c {
		fmt.Println(k, string(v))
	}


		statePopulations := make(map[string]int)
		statePopulations = map[string]int{
		"CA": 323423,
		"TX": 234234,
		"LA": 6456456,
		"NY": 19675,
	}

	// Prints just keys
	for k := range statePopulations {
		fmt.Println(k)
	}

	// Prints just values
	for _, v := range statePopulations {
		fmt.Println(v)
	}
}