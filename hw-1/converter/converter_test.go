package converter

import (
	"errors"
	"testing"
)

type convertorTest struct {
	in  interface{}
	out string
	err error
}

var convertorTests = []struct {
	name string
	in   interface{}
	out  string
	err  error
}{

	{"1", nil, "", ErrEmpty},
	{"2", 15.01, "15.01", nil},
	{"3", "some string", "some string", nil},
	{"4", true, "true", nil},
	{"5", -2147483649, "-2147483649", nil},
	{"6", int64(-2147483649), "-2147483649", nil},
}

func TestConvertToString(t *testing.T) {

	c := Converter{}

	for _, test := range convertorTests {
		t.Run(test.name, func(t *testing.T) {

			out, err := c.ConvertToString(test.in)
			if err != nil {
				if !errors.Is(err, test.err) {
					t.Errorf("ConvertToString() error = %v, wantErr %v", err, test.err)
				}
				return
			}
			if out != test.out {
				t.Errorf("ConvertToString() got != want\n%#v\n%#v", out, test.out)
			}
		})

	}
}
