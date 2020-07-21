package main

import "fmt"

type cat struct {
	name string
}

func (c cat) eat(food string) {
	println("Cat is eating: ", food)
}

type animal interface {
	eat(food string)
}

func sampleOneInterface() {
	var theCat animal = new(cat)
	theCat.eat("Fish")
}

type incrementer interface {
	increment() int
}

type theInt int

func (theInt *theInt) increment() int {
	*theInt++
	return int(*theInt)
}

func sampleTwoInterface() {
	var myInt theInt = theInt(0)
	var intPointer incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(intPointer.increment())
	}
}

func getCat(name string) cat {
	var _cat cat = cat{name: name}
	return _cat
}

func typeConversion() {
	println("==typeConversion()==")

	var myCat animal = getCat("Hori")
	otherCat, ok := myCat.(cat) //OR	otherCat   := myCat.(cat)
	if ok {
		otherCat.eat("Fish #2")
		fmt.Println(otherCat)
	} else {
		fmt.Println("Error conversion")
	}
}

func interfaceType() {
	var i interface{} = 0
	switch i.(type) {
	case int:
		println("Type Int")
	case string:
		println("Type String")
	}
}

func main() {

	sampleOneInterface()
	sampleTwoInterface()
	typeConversion()
	interfaceType()
}
