package objects

type PyString struct {
	Length int64
	Chars  []byte
}

func (this PyString) GetType() rune { return TYPE_STRING }

func (this PyString) String() string {
	return string(this.Chars)
}
