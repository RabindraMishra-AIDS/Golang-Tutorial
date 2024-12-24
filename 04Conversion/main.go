package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main()  {
	fmt.Println("Welcome to our pizza App")
	fmt.Println("Please rate us between 1 and 5")
	reader:=bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	fmt.Println("Rating received by user was:",input)

	// numRating=input + 1 [in this it will throw error because we input is of type string. ReadString was used.]

   //  Package used for string conversion in Go will be Strconv

   numRating,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
   if err!=nil{
	fmt.Println(err)
	//panic(err) can also be used to stop the program
   }
   fmt.Println("Added one to your rating:",numRating+1)
}