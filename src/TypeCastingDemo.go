package main

import (
	"fmt"
)

func typeconversion() {
	var x = 5
	var y float64 = float64(x)
	var z = int(y)
	var u float64 = 5.5
	fmt.Println(y, z, u)
}
