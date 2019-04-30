package main

import (
	"fmt"
)

func main () {
	//hys_if_declaration()
	//hys_infinit()
	//hys_enumerator()
	hys_enumerator_map()
}

func hys_if_declaration () {
	if i := 12; i > 50 {
		fmt.Println("over")
	} else {
		fmt.Println("less than")
	}
}

func hys_infinit () {
	i := 0
	for {
		i++
		if i > 1000000 {
			fmt.Println(i)
			break
		}
	}
}

func hys_enumerator_arr () {
	a := [...]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func hys_enumerator_map () {
	m := make(map[string]int) 
	m["user01"] = 1
	m["user02"] = 2
	m["user03"] = 3
	for k, v := range m {
		fmt.Println(k, v)
	}
}
