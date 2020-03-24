package main

import (
	"fmt"
	"reflect"
)

var (
	function reflect.Value
	inValue  []reflect.Value
	n        int
)

func bridge(funcPtr interface{}, args ...interface{}) {
	n = len(args)
	inValue = make([]reflect.Value, n)
	for i := 0; i < n; i++ {
		inValue[i] = reflect.ValueOf(args[i])
	}
	function = reflect.ValueOf(funcPtr)
	function.Call(inValue)
}

func callOne(v1 int, v2 int) {
	fmt.Println(v1, v2)
}

func callTwo(v1 int, v2 int, str string) {
	fmt.Println(v1, v2, str)
}

func main() {
	bridge(callOne, 1, 2)
	bridge(callTwo, 1, 2, "test2")
}
