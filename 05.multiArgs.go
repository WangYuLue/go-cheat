package main

import (
	"fmt"
	"reflect"
)

func sum(nums ...interface{}) {
	fmt.Println(reflect.TypeOf(nums[0]).String() == "int")
	fmt.Println(reflect.TypeOf(nums[1]).String() == "int")
	fmt.Println(reflect.TypeOf(nums[2]).String() == "int")
}

func main() {
	sum(1, 2, 3)
	sum("1", 2, []int{1, 2, 3})
	nums := []interface{}{1, 2, 3, 4}
	sum(nums...)
}
