package main

import "fmt"

func main() {
	fmt.Println("Understanding IF-Else in GoLang")

	amt_invested:=30000

	var result string

	if amt_invested >25000 { //Ensure that curly braces is written on this line only else the lexer will not be able to figure out and not  put semicolons automatically  thus need to follow golang internal structure
		result="Good invester"
	}else if amt_invested >20000{
		result="Regular investor"
	}else{
		result="Unreliable investor"
	}
	fmt.Println("Result:",result)

	if 4%2==0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	//In case of API handling we can use this syntax of initializing and checking the conditions
	if num:=4;num<10{
		fmt.Println("Number is Less than 10")
	}else{
		fmt.Println("Number is more than 10")
	}//This approach is  used to fetch api data and validate it at the sametime. (Uniqueness of GoLang)
}
