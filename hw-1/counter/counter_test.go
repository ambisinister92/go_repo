package counter_test

import (
	"testing"
  "reflect"

	. "hw-1/counter"
)

type countTest struct{
  in string
  num int
  out []string
}

var countTests = []countTest{

  {"|/*?%$^*()~!",4,[]string{"","","",""}},
  {"|/*?%$^*()~!",0,[]string{}},
  {"мама мыла Раму, Раму мыла мама !!!!",-1,nil},
  {"мама мыла Раму, Раму мыла мама !!!!",2,[]string{"мама : 2","мыла : 2"}},

}

func TestCount(t *testing.T){

  for _,test:= range countTests{
    out:=Count(test.in, test.num)
    if !reflect.DeepEqual(out,test.out) {
      t.Errorf("FormatUint(%v, %v) = %v want %v",
        test.in, test.num, out, test.out)
    }
  }

}
