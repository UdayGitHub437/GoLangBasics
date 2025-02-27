package main

import "fmt"

//assign values to slice and return it
func slice() {
	var s []string

	for i := 0; i < 5; i++ {
		s = append(s, fmt.Sprint(i))
	}
	fmt.Println(s[0])
}
