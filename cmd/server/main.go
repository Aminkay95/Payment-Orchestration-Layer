package main

import (
	"fmt"
	"net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API is up and running at port: ")
}

func main() {
	http.HandleFunc("/", helloWeb)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
