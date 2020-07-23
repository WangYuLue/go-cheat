package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("type1: ", reflect.TypeOf(num).Name())
	fmt.Println("type2: ", reflect.TypeOf(num).String())
	fmt.Println("value1: ", reflect.ValueOf(num))
	fmt.Println("value2: ", reflect.ValueOf(num).Type())
	fmt.Println("value3: ", reflect.ValueOf(num).Type().Name())
	fmt.Println("value4: ", reflect.ValueOf(num).Interface())
	fmt.Println("value5: ", reflect.ValueOf(num).Interface().(float64))
	val, ok := reflect.ValueOf(num).Interface().(float64)
	if ok {
		fmt.Println("val: ", val)
	}
}
