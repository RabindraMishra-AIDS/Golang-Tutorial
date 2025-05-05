package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to File Handling in GO")
	content := "I am on my way to Learn Golang. Add it to a File."

	//Using error ok syntax
	file, err := os.Create("./gofile.txt") //Creating  a File in current directory.
	if err != nil {
		panic(err)
	}
	// Writing the File
	length, err := io.WriteString(file, content) //returns Length
	checkerr(err)
	fmt.Println("Length of the file is:", length)

	//##################### Closing a File #######################
	defer file.Close() //Defer is used because it is possible that you may update the file further.

	//######################## Reading a File ###########################

	//Unlike file creation reading  does not requires os utility thus for this we use ioutil.
	readfile1("./gofile.txt")

}

func readfile1(filename string) {
	databyte, err := ioutil.ReadFile(filename) //whenever we read a file it is read in bytes and not in string format.
	checkerr(err)
	fmt.Println("Content After reading the file is:", string(databyte)) //Converting to string.
}

// Can also create a function to check error
func checkerr(err error) { //err has type error
	if err != nil {
		panic(err)
	}
}
