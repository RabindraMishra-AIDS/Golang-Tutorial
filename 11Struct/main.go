package main

import "fmt"

func main() {
	fmt.Println("Understanding structs in Go Lang")
	rabindra := User{"Rabindra","01tcet@gmail.com",20,1500000,true}
	fmt.Println("User struct:",rabindra)

	//for structure we must %+v format specifier
	fmt.Printf("\nDetails of rabindra:%+v",rabindra)
}

type User struct { //Declare first letter in Capital ie. Struct name "User" its attributes "Name,Email,etc"
	Name     string //Since if we would be exporting it then must follow the standard declarations
	Email    string
	Age      int
	Salary   int
	Verified bool
}
//Use printf while formating the the output strings