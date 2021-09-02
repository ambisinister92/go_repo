package converter

import (
	"errors"
	"strconv"
)

var ErrEmpty = errors.New("empty argument")

type ArgumentError struct {
	Func string
	Err  error
}

func (e *ArgumentError) Error() string {
	return e.Func + " : " + e.Err.Error()
}

func (e *ArgumentError) Unwrap() error { return e.Err }

func EmptyArgError(fn string) *ArgumentError {
	return &ArgumentError{fn, ErrEmpty}
}

type Converter struct {
}

func (c Converter) ConvertToString(value interface{}) (string, error) {
	const fnConvertToString = "ConvertToString"
	var s string

	if value == nil {
		return "", EmptyArgError(fnConvertToString)
	}

	switch value.(type) {
	case int, int64, int32, int16, int8:
		s = strconv.FormatInt(int64(value.(int64)), 10)
	case uint, uint64, uint32, uint16, uint8:
		s = strconv.FormatUint(uint64(value.(uint64)), 10)
	case bool:
		s = strconv.FormatBool(value.(bool))
	case float64, float32:
		s = strconv.FormatFloat(value.(float64), 'E', -1, 64)
	default:
		s = "unsupported value"
	}

	return s, nil
}
