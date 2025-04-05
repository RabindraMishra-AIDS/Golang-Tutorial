package main

import "fmt"

func main() {
	fmt.Println("Lets Learn Pointers") //Sometimes program creates a copy of variables thus we use pointers to ensure actually value is passed.

	var ptr *int 

	fmt.Println("The value of ptr is:",ptr) //Right now the value of ptr is none (nil).

	MyBtech_Rollno:=23
	var ptr1 = &MyBtech_Rollno

	fmt.Println("Address of MyBtech_Rollno is:",ptr1)
	fmt.Println("Value  that ptr1 is refrencing:",*ptr1)

	*ptr1=*ptr1+2
	fmt.Println("New ptr1 value is:",*ptr1) // Thus pointer ensure that operations are performed on actual values.

}