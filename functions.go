package main

import "fmt"

func addition(val1, val2 int) (result int) {
	result = val1 + val2
	return result
}

func substrac(value *int, substractVal int) {
	*value = *value - substractVal
}

func multipy(value1, value2 int) int {
	return value1 * value2
}

func function1() {
	println("==function #1==")
	a := 1
	b := 4
	result := addition(a, b)
	fmt.Println(a, "+", b, "=", result)
}

func sum(values ...int) *int {
	result := 0
	for i := 0; i < len(values); i++ {
		result += values[i]
	}
	return &result
}

func divide(value1, divideBy float64) (result float64, _err error) {
	if divideBy == 0.0 {
		///Creating ERROR Object///
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	result = value1 / divideBy
	return result, nil
}

func function3() {
	println("==function #2==")

	a := 10.0
	b := 0.0
	res, err := divide(a, b)

	if err != nil {
		println("Error:", err.Error())
	} else {
		fmt.Println("Result: ", res)
	}

}

func function2() {
	println("==function #2==")
	a := 4
	b := 1
	println("a:", a, "- b:", b)
	substrac(&a, b)
	println("a:", a)

}

func anonymousFunction() {
	println("== #anonymousFunction ==")
	func() {
		msg := "Hello World"
		fmt.Println("[anonymous func] ", msg)
	}() //==>Immediately INVOKED

	for i := 0; i < 10; i++ {
		func(value int) {
			fmt.Println("Current value: ", i)
		}(i)
	}

	var myMethod func() = func() {
		println("Hello world from anonymous function")
	}
	myMethod()

	var myMethodHasResult func(val string) (res string)
	myMethodHasResult = func(val string) string {
		return val + " __"
	}

	result := myMethodHasResult("HELLO")
	println("myMethodHasResult:", result)
}

func (animal *Animal) editName(name []string) string {

	for i := 0; i < len(name); i++ {
		animal.Name += " " + name[i]
	}
	return animal.Name
}

func sturctMethod() {
	println("====sturctMethod====")

	animal := new(Animal)
	animal.Name = "Cat"

	fmt.Println("ANIMAL:", animal)
	animal.editName([]string{"Catboy", "MAN"})
	fmt.Println("ANIMAL Name edited:", animal)
}

func main() {
	function1()
	function2()
	function3()
	anonymousFunction()
	sturctMethod()
}
