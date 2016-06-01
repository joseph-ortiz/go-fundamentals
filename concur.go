package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) //this makes the process parallel
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() { //the go keyword makes it a routine
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()
	waitGrp.Wait() //waits for all go routines to finish
}
