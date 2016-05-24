package main

import (
	"fmt"
)

func main() {

	//Delcare a slice called myCourses
	//slice of strings w/ length of 5 and a capacity of 10
	//mySlice := make([]int, 1, 4) //gets init with zero values
	mySlice := []int{1, 2, 3, 4, 5} //gets init with zero values
	//myCourses := []string{"Docker", "Puppet", "Python"}

	fmt.Println(mySlice)
	fmt.Printf("Length is %d. \nCapacity is: %d\n", len(mySlice), cap(mySlice))

	//for i := 1; i < 17; i++ {
	//	mySlice = append(mySlice, i)
	//	fmt.Printf("\nCapacity is: %d", cap(mySlice))
	//}
	for _, i := range mySlice {
		fmt.Println("for range loop:", i)
	}
	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...) //this appends all the elements to the slice
	fmt.Println(mySlice)
}
