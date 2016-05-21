package main

import (
	"fmt"
)

func main() {
	//Variables to store course rankings

	if firstRank, secondRank := 120, 614; firstRank < secondRank {
		fmt.Println("\nSecond course is doing better than first course")
		if firstRank > 100 {
			fmt.Println("Errr, you may wanna" +
				" consider another job :-D")
		}
	} else if firstRank > secondRank {
		fmt.Println("\nFirst course is doing better than second course")
	} else {
		fmt.Println("\nSomething weird happend here. We shouldn't have hit this line!")
	}
}
