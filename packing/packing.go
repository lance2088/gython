package packing

import (
	"bytes"
	"errors"
	"fmt"
)

var formatCharacters = map[rune][]interface{}{
	'h': []interface{}{nil, nil},
	'H': []interface{}{PackUnsignedShort, UnpackUnsignedShort},
	'c': []interface{}{PackChar, UnpackChar},
	'x': []interface{}{nil, nil},
	'b': []interface{}{nil, nil},
	'B': []interface{}{nil, nil},
	'i': []interface{}{nil, nil},
	'I': []interface{}{nil, nil},
	'l': []interface{}{nil, nil},
	'L': []interface{}{PackUnsignedLong, UnpackUnsignedLong},
	'?': []interface{}{nil, nil},
	'f': []interface{}{nil, nil},
	'd': []interface{}{nil, nil},
	's': []interface{}{nil, nil},
	'p': []interface{}{nil, nil},
}

func Pack(format string, args ...interface{}) ([]byte, error) {
	if len(format) != len(args) {
		errorString := "Format '%s' required %s arguments, %s provided"
		return nil, errors.New(fmt.Sprintf(errorString, format, len(format), len(args)))
	}
	var result bytes.Buffer
	var char rune
	var arg interface{}
	var err error

	for i := 0; i < len(format); i++ {
		char = rune(format[i])
		arg = args[i]

		funcs, ok := formatCharacters[char]
		if ok && funcs[0] != nil {
			err = funcs[0].(func(*bytes.Buffer, interface{}) error)(&result, arg)
		} else {
			err = errors.New(fmt.Sprintf("Unknown format character %s", char))
		}

		if err != nil {
			return nil, err
		}
	}

	return result.Bytes(), nil
}

func Unpack(format string, input []byte) ([]interface{}, error) {
	results := make([]interface{}, len(format))
	reader := bytes.NewReader(input)
	var char rune
	var err error
	var arg interface{}

	for i := 0; i < len(format); i++ {
		char = rune(format[i])

		funcs, ok := formatCharacters[char]
		if ok && funcs[1] != nil {
			arg, err = funcs[1].(func(*bytes.Reader) (interface{}, error))(reader)
		} else {
			err = errors.New(fmt.Sprintf("Unknown format character %s", char))
		}

		if err != nil {
			return nil, err
		}

		results[i] = arg
	}

	return results, nil
}
