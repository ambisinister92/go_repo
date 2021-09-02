package converter_test

import (
	"reflect"
	"testing"

	. "hw-1/converter"
)

type convertorTest struct {
	in  interface{}
	out string
	err error
}

var convertorTests = []convertorTest{

	{nil, "", ErrEmpty},
	{15.01, "15.01", nil},
	{"some string", "unsupported value", nil},
	{true, "true", nil},
	{-2147483649, "-2147483649", nil},
	{int64(-2147483649), "-2147483649", nil},
}

func TestConvertToString(t *testing.T) {

	for _,test := range convertorTests {
		c := Converter{}
		out, err := c.ConvertToString(test.in)
		if test.out != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("ConvertToString(%q) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}
