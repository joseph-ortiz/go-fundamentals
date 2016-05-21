package main

import (
	"fmt"
)

func main() {
	//Variables to store course rankings
	firstRank := "39"
	secondRank := "614"

	if firstRank < secondRank {
		fmt.Println("\nSecond course is doing better than first course")
	} else if firstRank > secondRank {
		fmt.Println("\nFirst course is doing better than second course")
	} else {
		fmt.Println("\nSomething weird happend here. We shouldn't have hit this line!")
	}
}
