package main

import "errors"

func returnError(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("j should not be zero")
	}
	var z = i / j
	return z, nil

}
