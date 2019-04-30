package main

import (
	"fmt"
)

func main () {
	//f = multi_return()
	//fmt.Println(f)
	//fmt.Println(arr_var())
	//var a [4]int
	fmt.Println(hys_sli_var())
}

func multi_return () ( one int , hello string) {
	one = 1
	hello = "hello"
	return one, hello
}

func hys_arr_var() ([3]int) { 
	var a [3]int
	return a
}

