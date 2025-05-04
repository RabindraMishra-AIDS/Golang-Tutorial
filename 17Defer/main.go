//"defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns,either because the surrounding function executed a return statement,reached the end of its function body,or because the corresponding goroutine is panicking.

package main

import "fmt"

func main() {
	defer fmt.Println("Hello ") //Here defer will result in printing hello at the last of the function based on LIFO if there are multiple defer
	defer fmt.Println("Priyansh Rathod")
	fmt.Println("My Name is Rabindra")

	myDefer()
}

func myDefer(){
	for i:=0;i<=3;i++{
		defer fmt.Println(i)
	}
}


//Understanding the flow of Execution
//Here myDefer() function will execute first because during function call they did not saw defer.
// My Name is Rabindra => myDefer() =>Priyansh Rathod => Hello
//for myDefer(): In stack it would be saved like [0,1,2,3] thus will print =>3,2,1,0
//Final Result:My name is Rabindra=>3,2,1,0=>Priyansh Rathod=>Hello