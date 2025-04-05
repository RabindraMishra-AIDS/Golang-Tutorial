package main
import (
	"fmt"
)
func main()  {
	fmt.Println("Welcome to Arrays in GOLANG")

	var myHero [3]string
	myHero[0]="Dr.APJ Abdul Kalam"
	myHero[2]="The Great Ashok Samrat"

	fmt.Println("myHero array:",myHero," ","Length of Array is:",len(myHero)) //Although there are just two elements length will show 3 elements and will include space in myHero Array

	var wishlist =[4]string{"AWS","Hadoop/Apache","Linux","RUST"}
	fmt.Println("MY Wishlist is:",wishlist)

}