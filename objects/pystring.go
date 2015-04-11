package objects

import "errors"

type PyString struct {
	Length int64
	Chars  []byte
}

func (this PyString) GetType() rune { return TYPE_STRING }

func (this PyString) String() string {
	return string(this.Chars)
}

func (this PyString) GetIndex(i int) (res byte, err error) {
	if int64(i) < this.Length {
		res = this.Chars[i]
	} else {
		err = errors.New("Index out of range")
	}

	return res, err
}
