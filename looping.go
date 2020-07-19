package main

import "fmt"

func simpleLoop() {
	println("==loop simple==")
	for i := 0; i < 10; i++ {
		fmt.Println("(i) Index: ", i)
	}

	println("-------------")

	j := 10
	for ; j < 20; j++ {
		fmt.Println("(j) Index: ", j)
	}
}

func whileLoop() {
	println("==while loop==")
	i := 0
	for i < 10 {
		fmt.Println("While i (", i, ") < 10 ")
		i++
	}
}

func infiniteLoop() {
	println("==infiniteLoop==")
	i := 5
loop:
	for {
		i++
		fmt.Println("i: ", i)
		if i > 15 {
			println("i more than 15")
			break loop //label not required
		}
	}
}

func loopCollection() {
	//ARRAY
	println("___array___")

	array := []string{"Jan", "Feb", "Mar", "Apr"}
	for key, value := range array {
		fmt.Println("key: ", key, " value: ", value)
	}

	println("____map key value___")

	//MAP
	theMap := map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4}
	for key, value := range theMap {
		fmt.Println("key: ", key, " value: ", value)
	}

	println("____map key ONLY")
	for key, _ := range theMap {
		fmt.Println("key: ", key, " value: -")
	}
}

func main() {
	println("_________LOOP_________")
	simpleLoop()
	whileLoop()
	infiniteLoop()
	loopCollection()
	println("_________(end) LOOP_________")
}
