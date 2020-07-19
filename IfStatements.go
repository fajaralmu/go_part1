package main

import "fmt"

func withMap() {
	months := map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
	}

	if monthIndex, ok := months["Jan"]; ok {
		println("Month index:", monthIndex)
	}

}

func switchKeyword() {
	println("==switch==")
	value := 2
	switch value {
	case 1, 2, 3:
		println("one until three")
		// automatically break
	case 4:
		println("four")

	case 5:
		println("five")

	default:

	}

	value2 := 7
	switch {
	case value2 >= 6 && value2 <= 10:
		println("value2: 6 to 10")
	case value2 > 10:
		println("value2 is bigger than 10")
	}
}

func typeSwitch() {
	println("==typeSwitch==")
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println(i, " type Integer")
	case float64:
		fmt.Println(i, " type float64")
	case string:
		fmt.Println(i, " type string")
	default:
		fmt.Println(i, " other type")
	}
}

func main() {

	if true {
		println("IF Statements")
		withMap()
		switchKeyword()
		typeSwitch()
	}
}
