package main

import (
	"fmt"

	"hw-1/converter"
	"hw-1/counter"
)


func main() {

	i := 15.01
	var j interface{}
	i64:=int64(i)
	y:=i64>0
	c:= converter.Converter{}
	fmt.Println(c.ConvertToString(i))
	fmt.Println(c.ConvertToString(j))
	fmt.Println(c.ConvertToString(i64))
	fmt.Println(c.ConvertToString(y))
	fmt.Println(counter.Count("Раз, два, три !!!!",3))

}
