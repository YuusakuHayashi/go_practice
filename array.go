package main

import (
	"fmt"
)

func main () {
	//fmt.Println(only_elements())
	//hys_slice_declaration()
	//hys_map_declaration()
	hys_map_el_exist()
}

func hys_only_elements() ([5]int) { 
	c := [...]int{1,2,3,4,5}
	return c
}

func hys_slice_basic() {
	arr := [...]int{1,2,3,4,5}

	sli := arr[0:3]
	//fmt.Println(arr)
	fmt.Println(sli)
}

func hys_slice_declaration() {
	sli := []int{1,2,3,4,5}
	fmt.Println(sli)
}

func hys_map_declaration() {
	m := map[string]int{"user01": 1, "user02": 2, "user03": 3}
	fmt.Println(m)
}

func hys_map_el_value() {
	m := map[string]int{"user01": 1, "user02": 2, "user03": 3}
	fmt.Println(m["user01"])
	fmt.Println(m["user02"])
	fmt.Println(m["user03"])
	fmt.Println(m["user04"])
}

func hys_map_el_exist() {
	m := map[string]int{"user01": 1, "user02": 2, "user03": 3}
	v, e := m["user01"]
	fmt.Println(v,e)
	v, e = m["user02"]
	fmt.Println(v,e)
	v, e = m["user03"]
	fmt.Println(v,e)
	v, e = m["user04"]
	fmt.Println(v,e)
}
