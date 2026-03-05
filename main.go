package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(Greet("World"))
	http.HandleFunc("/greet", greetHandler)
	http.ListenAndServe(":8080", nil)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprint(w, Greet(name))
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	return "Goodbye, " + name + "!"
}
