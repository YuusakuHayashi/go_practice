package main

import (
	"fmt"
)

type Hys struct {
	first string
	last  string
} 

func main () {
/*
	u0 := new(Hys)
	fmt.Println(u0)
	fmt.Println(*u0)

	var u Hys
	u.first = "hayashi"
	u.last  = "yuusaku"
	fmt.Println(u.first, u.last)

	u2 := new(Hys)
	u2.first = "hayashi"
	u2.last = "yuusaku"
	fmt.Println(u2.first, u2.last)
*/
	hys_structure_declaration ()
}

func hys_structure_declaration () {
	u := Hys{"hayashi", "yuusaku"}
	fmt.Println(u.first, u.last)
}
