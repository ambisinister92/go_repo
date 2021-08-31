package Human

import{

"fmt"
"math"

}

type Human interface{
  GetAge() int
}

type student struct{
  age int
}

func(s student) GetAge() int{
  return s.age
}

type teacher struct{
  age int
}
func(t teacher) GetAge() int{
  return t.age
}

type parent struct{
  age int
}
func(p parent) GetAge() int{
  return p.age
}

func CalculateAverageAge(h ...Humans) float64{
  
}
