package human

import "math"

type Human interface {
	GetAge() int
}

type Student struct {
	Age int
}

func (s Student) GetAge() int {
	return s.Age
}

type Teacher struct {
	Age int
}

func (t Teacher) GetAge() int {
	return t.Age
}

type Parent struct {
	Age int
}

func (p Parent) GetAge() int {
	return p.Age
}

func CalculateAverageAge(h ...Human) float64 {

	var allage float64

	for _, v := range h {
		allage += float64(v.GetAge())
	}

	intPart, fracPart := math.Modf(allage / float64(len(h)))
	fracPart= math.Round(fracPart * 100)/ 100

	return intPart + fracPart

}

