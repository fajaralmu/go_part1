package main

import (
	"fmt"
	"reflect"
)

func theMap() {
	println("===== map ====")
	months := make(map[string]int) //, 5)
	months = map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
		"May": 5,
	}
	printlns(months)
	months["Jun"] = 6
	delete(months, "Jan")
	println("map val:", months["Jun"])

	duplicate := months //Points to same data
	duplicate["Feb"] = 10
	println("FEB: ", months["Feb"])
}

type appUser struct {
	id       int
	username string
	name     string
}

func theStruct() {
	println("======= Struct =====")
	theUser := appUser{
		id:       1,
		username: "Fajar",
		name:     "Fajar ALM",
	}

	fmt.Println(theUser)
}

func anonymousStruct() {
	println("======anonymousStruct=======")
	visitor := struct{ name string }{name: "Sahal"}
	visitor2 := visitor //COPY DATA
	visitor2.name = "Achmed"
	fmt.Println(visitor)
	fmt.Println(visitor2)

	//Points to same data
	visitor3 := &visitor
	visitor3.name = "Toni"
	fmt.Println(visitor)
}

//Animal is a base type
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

//Bird is an animal
type Bird struct {
	Animal //embed
	canfly bool
	color  string
}

func inheritanceLike() {
	println("=====inheritanceLike======")
	bird := Bird{}
	bird.Name = "Eagle"
	bird.canfly = true
	bird.color = "black"

	printlns(bird)
}

func reflection() {
	printlns("=========reflection=======")
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	println("Name:", field.Name, " TAG:", field.Tag)
}

func main() {
	println("____ARRAY AND STRUCTS____")
	theMap()
	theStruct()
	anonymousStruct()
	inheritanceLike()
	reflection()
}
