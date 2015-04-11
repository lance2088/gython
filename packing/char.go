package packing

import (
	"bytes"
	"errors"
	"fmt"
)

func PackChar(buf *bytes.Buffer, arg interface{}) error {
	switch t := arg.(type) {
	case rune:
	default:
		return errors.New(fmt.Sprintf("Expected type rune, got %s", t))
	}

	buf.WriteRune(arg.(rune))
	return nil
}

func UnpackChar(reader *bytes.Reader) (interface{}, error) {
	ch, _, err := reader.ReadRune()
	return ch, err
}
