package main

import (
	"fmt"
	"sort"
)

func main() {
	//Note:In Slice we do not explicitly define array size thus size is expandable.
	//We need to intialize slice upon its creation either empty or no empty
	var fruitlist = []string{"Apple", "Tomato", "Pineapple"}
	fmt.Printf("The type of fruitlist is:%T\n",fruitlist) //Printf requires format specifier else will show error.

	//We can also append new data into Slice.
	fruitlist=append(fruitlist,"Mango","Lichi")
	fmt.Println(fruitlist)

	//Slicing 
	fmt.Println(fruitlist[1:]," ",fruitlist[:3])

	highscore:= make([]int, 4) //using make() keyword to initialize slice of size 4.
	highscore[0]=30
	highscore[1]=70
	highscore[2]=39
	highscore[3]=81

	fmt.Println("HighScores are",highscore)

	//Using Append in that case make() will reallocate the memory to accomodate all the elements.
   //However, if we normally write highscore[4]=90 it will give error due to size limit error (in this case 4 elements)	

   highscore=append(highscore,99,67 )
   fmt.Println("New HighScores are",highscore)


   //Unlike Arrays we can use 'sort' in slice also explore variour sort method by writing sort(.)
   sort.Ints(highscore)
   fmt.Println("Sorted Highscore:",highscore)

   //check whether sorted or not
   fmt.Println(sort.IntsAreSorted(highscore))
}