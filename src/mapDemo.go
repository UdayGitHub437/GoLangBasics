package main

import "fmt"

func mapD() {
	var y map[string]int
	y = make(map[string]int)
	y["kumar"] = 23
	fmt.Println(y["kumar"])
	x := make(map[string]int, 5)
	x["uday"] = 27
	x["virat"] = 33
	u, j := x["uday1"]
	fmt.Println(u, j)
	fmt.Print(x["uday"])
}
