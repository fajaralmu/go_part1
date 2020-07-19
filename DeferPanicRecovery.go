package main

import (
	"fmt"
	"net/http"
)

func deferFunction() {
	currenntHour := 15
	fmt.Println("Open The Door at ", currenntHour)
	//defer will be called last
	defer fmt.Println("Don't forget to Close the Door, enteting room at: ", currenntHour)
	fmt.Println("Enter the Room..")
	currenntHour += 2
	fmt.Println("Leaving the Room at: ", currenntHour)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"name\":\"Hello Go\"}"))

}

func panicFunction() {
	println("==panicFunction==")
	http.HandleFunc("/", handleMain)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Handle Error")
		panic(err.Error())
	}
}

func deferFunction2() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	//among defers, last defer wil be called 1st

}

func main_defer_panic() {
	println("____defer panic and recover___")
	deferFunction()
	deferFunction2()
	panicFunction()
}
