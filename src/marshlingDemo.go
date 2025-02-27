package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name         string
	Age          int
	MobileNumber int
}

func marshalStruct() ([]byte, error) {
	p := Person{
		Name:         "Uday",
		Age:          27,
		MobileNumber: 9908789159,
	}

	jsondata, err := json.Marshal(p)
	fmt.Println(string(jsondata))
	if err != nil {
		return nil, err
	}
	return jsondata, err

}
