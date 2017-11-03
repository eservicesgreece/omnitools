package main

import "reflect"

func checkIfIsMap(typeFC interface{}) bool {
	if reflect.ValueOf(typeFC).Kind() == reflect.Map {
		return true
	}
	return false
}
