package main

import "fmt"

func sample1() {
	println("======not pointer=====")
	a := 42
	b := a //not pointing, JUST COPY value
	fmt.Println("A:", a, " B:", b)
	a = 100 //B dose not change
	fmt.Println("A:", a, " B:", b)
}

func sample2() {
	println("======pointer=====")
	var a int = 100
	var b *int = &a //or b:=&a cannot add arithmatic

	fmt.Println("A:", a, " B(Memory address): ", b)
	//change from a
	a = 200
	//*B go to memory address  and pull the value
	fmt.Println("A:", a, " *B(Pull the value):", *b)

	//change from b
	*b = 350
	fmt.Println("A:", a, " *B(Currently changed):", *b)

}

func structPointer() {
	println("==structPointer==")
	var obj *object
	fmt.Println("*OBJ not initialized:", obj)
	obj = &object{id: 10}
	fmt.Println("OBJ initialized: ", obj)
	fmt.Println("(*)OBJ initialized: ", *obj)

	//assign value
	obj.id = 300
	fmt.Println("(*)OBJ property has been changed initialized: ", *obj)

}

type object struct {
	id int
}

func main() {

	println("___POINTER___")
	sample1()
	sample2()
	structPointer()
}
