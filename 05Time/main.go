package main

import ("fmt"
"time")

func main() {
	fmt.Println("Welcome to the Time package!")

	presentTime := time.Now() //Get current time

	fmt.Println("Current Time: ",presentTime)

	//We can also format time as per requirement
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //Note: This 01-02-2006 is unique way that GO follows to delcare a time format.

	//Creating a Date 
	createdDate:=time.Date(2021,time.April,10,23,0,0,0,time.UTC)
	fmt.Println("createdDate:",createdDate)
}
