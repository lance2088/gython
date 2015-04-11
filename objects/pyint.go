package objects

import "strconv"

type PyInt struct {
	Number int32
}

func (this PyInt) GetType() rune { return TYPE_INT }

func (this PyInt) String() string {
	return strconv.FormatInt(int64(this.Number), 10)
}
