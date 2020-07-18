package main

import "fmt"

func boolean() {
	println("=========BOOLEAN========")

	var n bool = true
	fmt.Printf("value: %v, type: %T\n", n, n)

}

func integer() {
	/**
	var i8 int8   //-128 > 127
	var i16 int16 // -32 768 > 32 767
	var i32 int32 //bigger
	var i64 int64 //bigger

	uint8 unit16, uint32 positives only
	**/
}

// func printValAndType(val, _type) {
// 	fmt.Printf("Value: %v, Type: %T \n", val, _type)
// }

func byteoperation() {
	println("====byteoperation=====")
	a := 10 //1010
	b := 3
	println(a & b)
	println(a | b)
	println(a ^ b)
	println(a &^ b)
}

func complexNumber() {
	println("====complexNumber=====")
	var a complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", (a), (a))
	fmt.Printf("%v, %T\n", real(a), real(a))

	var n complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", (n), (n))
}

func strings() {
	println("======== STRINGS =======")
	s := "this is String"
	firstChar := s[0]
	fmt.Printf("First Char %v, %T\n", firstChar, firstChar)

	//TO ASCII//
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

func runes() { //INT32
	println("==== RUNES =====")
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}

func numerics() {

	integer()
	byteoperation()
	complexNumber()

}

func main() {
	println("========PRIMITIVES=======")
	boolean()
	numerics()
	strings()
	runes()
}
