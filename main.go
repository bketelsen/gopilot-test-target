package main

import "fmt"

func main() {
	fmt.Println(Greet("World"))
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func Farewell(name string) string {
	return "Goodbye, " + name + "!"
}
