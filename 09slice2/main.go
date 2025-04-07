package main

import "fmt"

func main() {
	var highscore = []int{10, 45, 78, 62}
	fmt.Println("Highscore:", highscore)

	//Removing a Value from Slice based on index
	var programming = []string{"Python", "Ruby", "Swift", "flutter"}
	index := 2
	programming = append(programming[:index], programming[index+1:]...) //Last index is excluded in slicing.
	fmt.Println("Removed indexed 2:", programming)
	fmt.Println("Goconcept", goconcept())

}
