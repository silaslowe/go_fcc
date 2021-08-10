// package main

// import "fmt"

// func main() {
// 	statePopulations := make(map[string]int)
// 	statePopulations = map[string]int{
// 		"CA": 323423,
// 		"TX": 234234,
// 		"LA": 6456456,
// 		"NY": 19675,
// 	}
// 	// Add value
// 	statePopulations["GA"] = 100000123
// 	fmt.Println(statePopulations["GA"])
// 	// Delete value
// 	delete(statePopulations,"GA")
// 	fmt.Println(statePopulations)

// 	// ok determines if a value is in map
// 	pop := statePopulations["Oho"]
// 	fmt.Println(pop)

// 	pops, ok := statePopulations["Oho"]
// 	fmt.Println(pops, ok)

// 	// Length of map
// 	fmt.Println(len(statePopulations))

// 	sp := statePopulations
// 	delete(sp, "Ohio")
// 	fmt.Println(sp)
// 	fmt.Println(statePopulations)
// }