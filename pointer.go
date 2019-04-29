package main

import (
	"fmt"
)

func main () {
	var i0 int
	var i1 *int
	fmt.Println(i0)
	fmt.Println(i1)
	i1 = &i0 
	fmt.Println(i1)
	fmt.Println(*i1)
}
