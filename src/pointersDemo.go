package main

import (
	"fmt"
	"unsafe"
)

func pointer() {
	var i int = 437
	var j *int
	j = &i
	fmt.Println(*j)
}

type animal struct {
	name   string
	colour string
}

func initializeStruct() {
	a := animal{
		name:   "dog",
		colour: "black",
	}
	fmt.Print(a)
	a = animal{
		name:   "horse",
		colour: "white",
	}
	fmt.Println(a)
}

type month struct {
	name string
	days int
}

func fun1(m1 *month) {
	m := month{
		name: "March",
		days: 31,
	}
	fmt.Println(m)
}

func fun3() {
	m := month{
		name: "January",
		days: 31,
	}
	fun1(&m)
	fmt.Println(m.name)
}

func fun4() {
	var num int = 100

	// Get the address of the variable
	address := unsafe.Pointer(&num)

	// Convert the pointer to *int and dereference to get the value
	ptr := (*int)(address)
	fmt.Println("Value at specific address:", *ptr)
}
