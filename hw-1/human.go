package main

import "math"
import "fmt"

type Human interface {
	GetAge() int
}

type student struct {
	age int
}

func (s student) GetAge() int {
	return s.age
}

type teacher struct {
	age int
}

func (t teacher) GetAge() int {
	return t.age
}

type parent struct {
	age int
}

func (p parent) GetAge() int {
	return p.age
}

func CalculateAverageAge(h ...Human) float64 {

	var allage float64

	for _, v := range h {
		allage += float64(v.GetAge())
	}

	int, frac := math.Modf(allage / float64(len(h)))
	frac, _ = math.Modf(frac * 100)
	frac = frac / 100

	return int + frac

}

func main() {

	t1 := teacher{24}
	s1 := student{16}
	p1 := parent{48}
	fmt.Println(t1.GetAge())
	fmt.Println(s1.GetAge())
	fmt.Println(p1.GetAge())
	fmt.Println(CalculateAverageAge(t1, s1, p1))

}
