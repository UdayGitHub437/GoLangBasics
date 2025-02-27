package main

import "fmt"

var ei interface{}

func emptyInterface() {
	ei = 437
	fmt.Println(ei)
	ei = "Uday"
	fmt.Println(ei)
	ei = 99.99
	fmt.Println(ei)
}
