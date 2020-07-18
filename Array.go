package main

import "fmt"

func array() {
	println("=======ARRAY1=======")

	grades := [3]int{90, 89, 70}
	fmt.Printf("Grades 1: %v \n", grades)
	grades2 := [...]int{90, 89, 70}
	fmt.Printf("Grades 2: %v \n", grades2)

	var students [3]string
	students[0] = "Ahmed"
	students[1] = "Hasan"
	students[2] = "Husen"

	printValAndTypeForSingleVal(students)
	fmt.Printf("Array Length: %v\n", len(students))
}

func arrayPointing() {
	println("=======arrayPointing======")
	a := [...]int{1, 2, 3}
	b := &a    //pointing same data
	b[1] = 100 //a value changed
	printValAndTypeForSingleVal(a)
	printValAndTypeForSingleVal(b)
}

func slice() {
	println("=======slice=========")
	a := []int{1, 2, 3}
	b := a     //SLICE ALWAYS pointing same data
	b[1] = 100 //a value changed
	printValAndTypeForSingleVal(a)
	printValAndTypeForSingleVal(b)
	fmt.Printf("Array Length: %v\n", len(a))

}

func slice2() {
	println("========slice2=======")
	///All these variables point to same data///
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //slice all of a
	c := a[3:]  //slice from index 3 to end
	d := a[:6]  //slice from index 0 to 6
	e := a[3:6] //slice from index 3 to 6
	printlns(a, b, c, d, e)

}

func sliceManipulation() {
	println("======MAkingSliCE====")
	a := make([]int, 3)
	println(a)
	println("length: ", len(a))
	println("capacity: ", cap(a))

	println("__Slice APPEND__")
	//must append same type
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3, 4, 5, 6)
	println("length: ", len(a))
	println("capacity: ", cap(a))

	b := append(a[3:], a[4:]...)
	fmt.Println(b)
}

func printlns(val ...interface{}) {
	println("array length: ", len(val))
	for i := 0; i < len(val); i++ {
		fmt.Println(val[i])
	}
}

func main() {
	println("________ARRAY_______")
	array()
	arrayPointing()
	slice()
	slice2()
	sliceManipulation()
}
