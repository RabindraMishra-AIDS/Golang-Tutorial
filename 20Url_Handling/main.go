package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.isro.gov.in/"

func main() {
	fmt.Println("Handling URL in Go")

	//parsing the URL
	result, _ := url.Parse(myurl)
	fmt.Println(result)

	//Extracting info  
	fmt.Println(result.Scheme) 
	fmt.Println(result.Host) 
	fmt.Println(result.Path) 
	fmt.Println(result.RawQuery) 
	fmt.Println(result.Port()) 

	queryParameters:=result.Query() //Query() Will store all the parameters of url in more better way.

	fmt.Printf("The type of query params are:%T\n",queryParameters)

	for _,val:= range queryParameters{
		fmt.Println("Param is:",val)
	} 

	partsOfUrl:=&url.URL{ //Here we need to Pass Reference instead of actual copy
		Scheme:"https",
		Host:"www.isro.gov.in",
		Path:"/",
		RawPath:"Rabindra",

	}
	anotherURL:=partsOfUrl.String()
	fmt.Println(anotherURL)
}
