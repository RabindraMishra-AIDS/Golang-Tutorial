package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.jagranjosh.com/general-knowledge/top-famous-freedom-fighters-of-india-1737527201-1"

func main() {
	fmt.Println("Web Request in GoLang")
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response of our Get Method is:%T\n", req)
	defer req.Body.Close() //Closing the connection

	//########## Reading the Request using ioutil ###############
	//Will be using readall since it can read all format files,directory,etc.

	databyte, err := io.ReadAll(req.Body) //Can also write req.Body
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of req.Header is:", string(databyte))
}
