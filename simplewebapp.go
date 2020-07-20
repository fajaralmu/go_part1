package main

import (
	"fmt"
	"net/http"
)

//PORT is portss
const PORT string = ":8011"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bird Bird = Bird{
		Animal: Animal{
			Name:   "eagle",
			Origin: "Java",
		},
		color: "black",
	}
	bird.canfly = false
	//json.NewEncoder(w).Encode(bird)
	w.Write([]byte("{\"name\":\"Hello Go\"}"))
	fmt.Println("handleIndex END..")
}

func serve() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Handle Error")
		panic(err.Error())
	}

}

func main_web() {
	serve()
}
