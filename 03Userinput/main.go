package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome our dear Users"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) //Note we are reading from os(Stdin) ie.StandardInput output.
	fmt.Println("Enter the rating of the services: ")

	//Comma ok|| error
	//if error is there then it will get store in '_' varibale and output may get store in to input variable
	//simply think of this that first part is try varibale and second part is catch variable.
	input,_:=reader.ReadString('\n')
	fmt.Println("Rating for our given service is",input)
	fmt.Printf("Tye of Rating is:%T",input)  //But there is one problem here it is showing type of input is string.

	//Note: If we are not concerned with input then instead of input we can also write like '_'.
}
