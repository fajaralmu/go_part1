package main

import (
	"fmt"
	"sync"
	"time"
)

func greet() {
	fmt.Println("Good Morning")
	println("go func")
	//	time.Sleep(100 * time.Millisecond)

}
func gRSampleOne() {

	println("===#1===")
	go greet()

	println("___________")
	time.Sleep(100 * time.Millisecond)
}

var wg = sync.WaitGroup{}

func gRSampleTwo() {
	println("===#2===")

	wg.Add(1)
	msg := "Hellow"
	go func(msg string) {
		fmt.Println("MSG:", msg)
		wg.Done()
	}(msg)
	msg = "Hellaw"
	println("___________")
	wg.Wait()
}

func main() {

	gRSampleOne()
	gRSampleTwo()
}
