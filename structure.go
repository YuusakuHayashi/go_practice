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
	//hys_structure_declaration ()

	//u := Hys{"hayashi", "yuusaku"}
	//u.hys_method()
	h := Hys{"hayashi", "yuusaku"}
	h.hys_method()

	//y := Ysk{"hayashi", "yuusaku"}
	y := Ysk{}
	y.hys_method()
	hys_variable_type ("hello")
	hys_variable_type (23)
	hys_type_assertion("hello")
	hys_type_assertion(23)
*/
	hys_type_switch("hello")
	hys_type_switch(23)
}

func hys_structure_declaration () {
	u := Hys{"hayashi", "yuusaku"}
	fmt.Println(u.first, u.last)
}

func hys_variable_type (any interface{}) {
	fmt.Println(any)
}

func hys_type_assertion (any interface{}) {
	v, b := any.(string)
	println(v, b)
}

func hys_type_switch (any interface{}) {
	switch any.(type) {
		case string:
			fmt.Println("type is string =", any)
		case int:
			fmt.Println("type is integer =", any)
		default:
			fmt.Println("type is others =", any)
	}
}
