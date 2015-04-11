package utils

import "bytes"

func ReadBytes(reader *bytes.Reader, length int) ([]byte, error) {
	var next = make([]byte, length)

	var char byte
	var err error
	for i := 0; i < length; i++ {
		char, err = reader.ReadByte()
		if err != nil {
			return nil, err
		}
		next[i] = char
	}

	return next, nil
}
