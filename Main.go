package main

import (
	"fmt"
	"strconv"
)

//var CAPITAL string = "EXPOSED every where"
var global float32 = 50 //lowercase only exposed to the same package

func variable1() {
	fmt.Println("======vFUNC: Variable1======")

	var i int
	i = 42
	i = 27
	fmt.Println("Variable ", i)

	var j int = 10
	fmt.Println("Variable initialized", j)
	k := 45
	fmt.Println("Variable initialized 2: ", k)

	fmt.Printf("PrintF VALUE: %v, TYPE: %T \n", global, global)
}

func variable2() {
	println("=======VARIABLE_2========")

	var (
		name    string = "FAJAR"
		address string = "KEBUMEN"
		myURL   string = "github.com/fajaralmu" //ACRONIM MUST BE UPPERCASE
	)

	println("MY name is ", name, " i am from: ", address, " \n git repo: ", myURL)
}

func variableConversion() {
	println("==========variableConversion============")
	var i int = 20
	fmt.Printf("value: %v, type: %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("int to FLOAT value: %v, type: %T\n", j, j)

	var s string
	s = strconv.Itoa(i)
	fmt.Printf("int to String value: %v, type: %T\n", s, s)

}

func _main() {
	variable1()
	variable2()
	variableConversion()
}
