package main

import "fmt"

func forEach() {
	mp := make(map[string]int)
	mp = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for _, x := range mp {

		fmt.Println(x)
	}
}
