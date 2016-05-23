package main

import (
	"fmt"
)

func main() {

	//Delcare a slice called myCourses
	//slice of strings w/ length of 5 and a capacity of 10
	//	myCourses := make([]string, 5, 10) //gets init with zero values
	myCourses := []string{"Docker", "Puppet", "Python"}

	fmt.Printf("Length is %d. \nCapacity is: %d", len(myCourses), cap(myCourses))
}
