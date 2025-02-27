package main

import (
	"fmt"
	"reflect"
)

var Name = "Kohli"

func fun() {
	mp := make(map[string]int)
	mp = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	m := reflect.ValueOf(mp)
	mp["B"] = 6
	fmt.Println(m, mp)
}

func fun2(mp map[string]int) {
	mp["A"] = 5
}

func callErrorReturnFunc() {
	div, err := returnError(4, 0)
	if err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Print(div)
	}
}

func main() {
	typeAssertion()
	// emptyInterface()
	// g := rect{
	// 	height: 44,
	// 	width:  32,
	// }
	// r := rect{
	// 	height: 4.5,
	// 	width:  6.2,
	// }
	// c := circle{
	// 	radius: 3.2,
	// }
	// measure(g)
	// measure(r)
	// measure(c)
	// m := square{
	// 	side: 4.5,
	// }
	// fmt.Println(m.area())
	// fmt.Println(m.perim())
	// var u int = 437
	// passByValue(u)
	// // fmt.Println(u)
	// passByReference(&u)
	// fmt.Println(u)
	// fun3()
	// initializeStruct()
	// pointer()
	// fmt.Println(recursion(5))
	// variadicfunc(3, 4, 5, 6, 7, 8)
	// x := []int{1, 2, 3}
	// variadicfunc(x...)
	// mapD()
	// slice()
	// array(5)
	// jsondata, _ := marshalStruct()
	// fmt.Println(string(jsondata))
	// fmt.Println(unmarshling())
	// callErrorReturnFunc()
	// typeconversion()
	// forEach()
	// fun()
	// temp1(4)
}

func temp1(x int) {
	fmt.Println(x)
	temp2(x)
	fmt.Println(x)
}

func temp2(x int) {
	x = 6
	fmt.Println(x)
}
