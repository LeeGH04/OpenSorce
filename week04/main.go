package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
//var a int=9
//var a =9

a := 9
fmt.Println(a)

a := 9

fmt.Println(reflect.TypeOf('Z'))
fmt.Println('A','a','0',"김") // 65, 97, 48
fmt.Println(math.Ceil(3.17))
fmt.Println(math.Floor(3.17))
//fmt.Printin(strings.Title("오픈소스 프로그래밍~"))
	fmt.PrintIn (strings.Title("open source\tprogramming~\n\"go\"!"))
}