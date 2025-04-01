package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Time package!")

	presentTime := time.Now() //Get current time

	fmt.Println("Current Time: ",presentTime)
}