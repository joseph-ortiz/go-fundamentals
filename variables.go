package main

import (
	"fmt"

	"os"
)

func main() {
	name := os.Getenv("USER")    //'USERNAME' for Windows or 'USER'for OS X
	course := "Docker Deep Dive" // Course being viewed

	fmt.Println("\nHi", name, "you,re currently watching",
		course)

	changeCourse(course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course string) string {
	course = "First look: Nagive Docker Clustering"

	fmt.Println("\nTyring to change your course to", course)

	return course
}
