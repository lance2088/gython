package packing

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/brettlangdon/gython/packing/utils"
)

func PackUnsignedLong(buf *bytes.Buffer, arg interface{}) error {
	switch t := arg.(type) {
	case int64:
	default:
		return errors.New(fmt.Sprintf("Expected type int64, got %s", t))
	}

	num := arg.(int64)
	buf.WriteByte(byte((num & 0xff)))
	buf.WriteByte(byte(((num >> 8) & 0xff)))
	buf.WriteByte(byte(((num >> 16) & 0xff)))
	buf.WriteByte(byte(((num >> 24) & 0xff)))

	return nil
}

func UnpackUnsignedLong(reader *bytes.Reader) (interface{}, error) {
	var result int64

	next, err := utils.ReadBytes(reader, 4)

	if err != nil {
		return nil, err
	}

	result = int64(next[0])
	result |= int64(next[1]) << 8
	result |= int64(next[2]) << 16
	result |= int64(next[3]) << 24

	return result, nil
}
