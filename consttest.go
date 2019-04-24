package main

import (
	"fmt"
)

const weekday = "weekday"

func main () {
	const (
		sun = iota
		mon
		tue
		wed
		thr
		fri
		sat
	)
	fmt.Println(weekday)
	fmt.Println(sun)
	fmt.Println(mon)
	fmt.Println(tue)
	fmt.Println(wed)
	fmt.Println(thr)
	fmt.Println(fri)
	fmt.Println(sat)
}
