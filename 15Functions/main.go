package main

import "fmt"

func main() {
	greetagain()
	fmt.Println("Welcome to function is golang")
	greet() //Function Call
	result:=adder(3,5)
	fmt.Println("Addition of result is:",result)
	proresult:=proAdder(7,8,9,6,8)
	fmt.Println("proResult is:",proresult)
	i,j :=special()
	fmt.Println("i,j",i,j)

}
func greet(){
	fmt.Println("Namastey From Rabindra")
}

func greetagain(){
	fmt.Println("Again Hello")

}

//Note:- In Golang defining function inside function is not supportive.
func adder(temp1 int,temp2 int) int { //Function signature that returns int
	return temp1+temp2
}

//getting unkown number of parameters.
func proAdder(value ...int)int{
	total:=0
	for _,val:=range value{
		total+=val
	}
	return total
}
// We can return two values also in a function 
func special()(int,string){ // will return int and string
	return 10,"Hello"

}