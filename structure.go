package main

import (
	"fmt"
)

type hys_interface interface {
	hys_method()
}

type Hys struct {
	first string
	last  string
} 

type Ysk struct {
	//first string
	//last  string
} 

func (h Hys) hys_method () {
	//fmt.Printf("hello my name is %s %s \n",h.first, h.last)
	fmt.Printf("hello\n")
}

func (y Ysk) hys_method () {
	//fmt.Printf("hello my name is %s %s \n",y.first, y.last)
	fmt.Printf("hello\n")
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
	//hys_structure_declaration ()

	//u := Hys{"hayashi", "yuusaku"}
	//u.hys_method()
	h := Hys{"hayashi", "yuusaku"}
	h.hys_method()

	//y := Ysk{"hayashi", "yuusaku"}
	y := Ysk{}
	y.hys_method()
}

func hys_structure_declaration () {
	u := Hys{"hayashi", "yuusaku"}
	fmt.Println(u.first, u.last)
}

