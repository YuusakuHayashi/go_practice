package main

import (
	"fmt"
)

func main () {
	fmt.Println(named_return())
}

func named_return () (hello string) {
	hello = "hello world"
	return hello 
}
