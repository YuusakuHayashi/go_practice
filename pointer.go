package main

import (
	"fmt"
)

func main () {
	var i0 int
	var i1 *int
	fmt.Println(i1)		// 初期値は<nil>
	i1 = &i0 
	fmt.Println(i1)		//i0のアドレスを出力
	fmt.Println(*i1)	//i0の値を参照
	//i1 = i0 		//コンパイルエラー
}
