package main

//Best practives
// #1 Don't create goroutines in libraries
//  - Let cuntmer control concurrency
// #2 WHen creating a goroutine, know how it will end
//  - avids subtle memory leaks
// #3 Check for race conditions at compile time
//  - check race condiiton: command> go run race ...

import (
	"container/list"
	"fmt"
	"runtime"
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
var mutex = sync.RWMutex{}
var counter int = 0

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

func gRSampleThree() {
	println("________gRSampleThree_______")

	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go printHello()
		mutex.Lock()
		go updateCounter()
	}

	wg.Wait()
}

func printHello() {

	fmt.Println("Hello world ", counter)
	mutex.RUnlock()
	wg.Done()

}

func updateCounter() {

	counter++
	mutex.Unlock()
	wg.Done()

}

func sampleMaxProcs() {
	println("_____-sampleMaxProcs______")

	//GOMAXPROCS return peviously set thread
	//negative values does not changing anything
	fmt.Println("THreads: ", runtime.GOMAXPROCS(-1))
}

func theList() {
	println("___theList___")

	var myList *list.List = list.New()
	myList.PushBack("1")
	myList.PushBack("2")
	myList.PushBack("3")

	fmt.Println("List size: ", myList.Len())

	println("__back - front__")
	for e := myList.Back(); e != nil; e = e.Prev() {
		fmt.Println("Value: ", e.Value)
	}
	println("__front - back__")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println("Value: ", e.Value)
	}
}

func main_goRoutine() {

	gRSampleOne()
	gRSampleTwo()
	gRSampleThree()
	sampleMaxProcs()
	theList()
}
