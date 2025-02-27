package main

import (
	"encoding/json"
	"fmt"
)

func unmarshling() error {
	jsondata, _ := marshalStruct()
	type Person1 struct {
		Name         string `json:"Name"`
		Age          int
		NobileNumber string
	}
	var p1 Person1
	fmt.Println(string(jsondata))
	err := json.Unmarshal(jsondata, &p1)
	if err != nil {
		return err
	}
	fmt.Println(p1.Name, p1.Age, p1.NobileNumber)
	return nil
}
