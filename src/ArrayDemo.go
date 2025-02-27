package main

import "fmt"

var x = [...]int{4, 5, 6}

func array(n int) {
	x := make([]int, n)
	fmt.Println(x[2])
	fmt.Println("length of array is", len(x))
}
