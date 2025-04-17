package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in GOLang")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} //Using Slice Data Structure.
	fmt.Println(days)

	//range keyword in Golanf automatically loops through array or Slice.
	for i := range days {
		fmt.Println(days[i]) //Here unlike JS i will return index instead of actual value.
	}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	//For Each type Loop.
	for index, value := range days {
		fmt.Printf("Index is:%v|| value is: %v\n", index, value)
	}

	//Funfact while Loop in Golang.🤨
	loopvalue := 5
	for loopvalue >= 0 {

		if loopvalue == 2 {
			loopvalue--
			continue
		}
		if loopvalue == 1 {
			loopvalue--
			goto rm

		}

		fmt.Printf("Value is: %v\n", loopvalue)
		loopvalue--
	}

	//Continue : Simply skips that Particular iteration.
   //Break :Terminates the Loop.

rm:
	fmt.Println("Welcome to world of Golang")

	temp := 3
	for temp >= 0 {
		fmt.Println("Value:", temp)
		temp--
	}

}
