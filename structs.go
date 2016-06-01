package main

import "fmt"

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//// You can instantiate a struct this way
	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	//composite literal
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println("\nDocker Deep Dive author is: ", DockerDeepDive.Author)
}
