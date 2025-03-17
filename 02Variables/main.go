package main

import "fmt"
//jwt:=3000 this is not allowed beacuse walrus operator is only allowed inside a method & cannot be used in global variable.
const Id ="12euo"   //Now it has become as a public variable becasue "I" is capital.
func main() {
	var username string = "Rabindra Mishra"
	fmt.Println(username)
	fmt.Printf("Variable username is of type: %T \n",username) //Here %T is a placeholder \n (denotes new Line).

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable username is of type: %T \n",isLoggedin)

	var smallval uint8 = 255   //Max range of smallval will be 255 (0-255)
	fmt.Println(smallval)
	fmt.Printf("Variable username is of type: %T \n",smallval)

	var value int= 999999  
	fmt.Println(value)
	fmt.Printf("Variable username is of type: %T \n",value)

	var price float64= 10.26   //Can also use float32 
	fmt.Println(price)
	fmt.Printf("Variable username is of type: %T \n",price)

	// Defaut values and some Aliases
	var rabindra float64   //In this case what be the default float value ? will it be any garbage value? 
	fmt.Println(rabindra)
	fmt.Printf("Variable username is of type: %T \n",rabindra)



	// Implicit way to declare a variable 

	var webiste="Learning.in"   //This also one way to declare a variable here the lexer will do the work and assume it to be as string.
	fmt.Println(webiste)


	// no var style of declaring a varibale using walrus operator
	numberofuser:=300                    //You can simply change the value of varibale in this type
	fmt.Println(numberofuser)
	numberofuser=2

	fmt.Println(Id)  //printing public variable.
}