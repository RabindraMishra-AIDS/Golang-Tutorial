package main

import "fmt"

//Also kanown as Hash/Ket value pairs in some other Languages.
func main() {
	fmt.Println("Understanding Maps in GO")

	//new gives a Lot of errors and is zeroed ie. cannot initialize thus we mostly use make()
	languages:= make(map[int]string) //Here I am defining that my key will be of type int and value will be of type string.
	languages[0]="Python"
	languages[1]="Ruby"
	languages[2]="Rust"
	languages[3]="PHP"

	fmt.Println("List of Languages:",languages)
	//Acessing value using key
	fmt.Println("Accesing key 2:",languages[2])

	//Deleting a key value pair
	delete(languages,2)

	//for loop in Go
	for i,j := range languages{ //Note: Walrus operator takes care of comma ok syntax
		fmt.Printf("For key %v -> value is %v\n",i,j) 
	} 
}
