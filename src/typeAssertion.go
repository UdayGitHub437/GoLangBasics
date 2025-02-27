package main

import "fmt"

func typeAssertion() {
	var ju interface{}
	ju = 4
	x, ok := ju.(string)
	if ok {
		fmt.Println(x)
	} else {
		panic(!ok)
	}
}
