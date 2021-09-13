package main

import (
	"fmt"

	"hw-1/converter"
	"hw-1/counter"
	"hw-1/human"
)

func main() {

	i := 15.01
	var j interface{}
	i64 := int64(i)
	y := i64 > 0
	c := converter.Converter{}
	fmt.Println(c.ConvertToString(i))
	fmt.Println(c.ConvertToString(j))
	fmt.Println(c.ConvertToString(i64))
	fmt.Println(c.ConvertToString(y))
	fmt.Println(counter.Count("Раз, два, три !!!!", 3))
	t1 := human.Teacher{24}
	s1 := human.Student{16}
	p1 := human.Parent{49}
	fmt.Println(t1.GetAge())
	fmt.Println(s1.GetAge())
	fmt.Println(p1.GetAge())
	fmt.Println(human.CalculateAverageAge(t1, s1, p1))

}
