package packing

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/brettlangdon/gython/packing/utils"
)

func PackUnsignedShort(buf *bytes.Buffer, arg interface{}) error {
	switch t := arg.(type) {
	case int, int32:
	default:
		return errors.New(fmt.Sprintf("Expected type int32, got %s", t))
	}

	buf.WriteByte(byte((arg.(int) & 0xff)))
	buf.WriteByte(byte(((arg.(int) >> 8) & 0xff)))

	return nil
}

func UnpackUnsignedShort(reader *bytes.Reader) (interface{}, error) {
	next, err := utils.ReadBytes(reader, 2)
	if err != nil {
		return 0, err
	}

	return int32(((int(next[1]) & 0xFF) << 8) | (int(next[0]) & 0xFF)), nil
}
