package main

import "fmt"

func main() {
	name := "bill"
	namePointer := &name
	fmt.Println("A: ", namePointer)
	fmt.Println("B: ", &namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println("C: ", namePointer)
	fmt.Println("D: ", &namePointer)
}
