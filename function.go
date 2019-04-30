package main

import (
	"fmt"
)

func main () {
	//lumbda_function()
	//fmt.Println(array_return())
}

func hys_named_return () (hello string) {
	hello = "hello"
	return hello
}

func hys_multi_return () ( one int , hello string) {
	one = 1
	hello = "hello"
	return one, hello
}

func hys_anonymous () {
	func (hello string, world string) {
		fmt.Println(hello + " " + world)
	}("hello", "world")

}

func hys_array_return () ([3]int) { 
	var a [3]int
	return a
}

func hys_slice_return () ([]int) { 
	a := [...]int{1,2,3}
	s := a[:3]
	return s 
}

func hys_map_return() (map[string]int) {
	m := make(map[string]int)
	return m
}
