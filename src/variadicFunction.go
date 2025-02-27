package main

import "fmt"

func variadicfunc(x ...int) {
	var z int
	for _, y1 := range x {
		z += y1
		fmt.Println(z)
	}

}
