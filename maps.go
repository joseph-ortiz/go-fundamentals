package main

import "fmt"

func main() {

	//create maps like this
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["NewCastle"] = 4

	// or like this
	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n", leagueTitles, recentHead2Head)
}
