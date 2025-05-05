// Functions wrapped up in Classes are called as Methods.
package main

import "fmt"

func main() {
	fmt.Println("Understanding structs in Go Lang")
	rabindra := User{"Rabindra","01tcet@gmail.com",20,1500000,true}
	fmt.Println("User struct:",rabindra)

	//for structure we must %+v format specifier
	fmt.Printf("\nDetails of rabindra:%+v\n",rabindra)

	//################## Methods #########################
	rabindra.GetStatus()
	rabindra.NewEmail("Golang@fun.com") //is the Value changed permanently?

	fmt.Printf("\nDetails of rabindra:%+v\n",rabindra)
}

type User struct { //Declare first letter in Capital ie. Struct name "User" its attributes "Name,Email,etc"
	Name     string //Since if we would be exporting it then must follow the standard declarations
	Email    string
	Age      int
	Salary   int
	Verified bool
}
//Use printf while formating the the output strings

func(rab User) GetStatus(){ //Here rab is a type of User. GetStatus is Method Name
	fmt.Println("Is User Verified?=>",rab.Verified) //If we want to export the Method First letter must be capital here it is "GetStatus"

}
func(u User) NewEmail(m string){
u.Email=m 
fmt.Println("New Email of user is:",u.Email)

}

//Note:Need to pass reference of Actual Object else Methods will overwrite on copies of objects and not on actual objects