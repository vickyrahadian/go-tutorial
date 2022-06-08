package main

import "fmt"

func main() {

	message := "Hello All"
	var m *string
	m = &message

	fmt.Println(m)
	message = "Test"
	fmt.Println(m)
}
