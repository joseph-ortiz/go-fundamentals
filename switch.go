package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	basicSwitch()
	randomNumberSwitch()
}

func basicSwitch() {

	switch topic := "docker"; topic {
	case "linux":
		fmt.Println("\nHere are some recommended Linux courses")
	case "docker":
		fmt.Println("\nHere are some recommended Docker courses")
		fallthrough
	case "windows":
		fmt.Println("\nHere are some recommended Windows courses")
	default:
		fmt.Println("\nHere are some recommended Linux courses")
	}
}

func randomNumberSwitch() {
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("we got an even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("we got an odd number -", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
