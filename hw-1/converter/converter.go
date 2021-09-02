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

	if value == nil {
		return "", EmptyArgError(fnConvertToString)
	}

	switch s := value.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	default:
		return "unsupported value", nil
	}

}
