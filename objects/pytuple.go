package objects

import "strings"

type PyTuple struct {
	Length int64
	items  []PyObject
}

func (this PyTuple) GetType() rune { return TYPE_TUPLE }

func (this PyTuple) String() string {
	output := []string{}
	for _, value := range this.items {
		output = append(output, value.String())
	}
	return "(" + strings.Join(output, ", ") + ")"
}

func (this PyTuple) SetItem(i int, obj PyObject) {
	this.items[i] = obj
}

func (this PyTuple) GetItem(i int) PyObject {
	if int64(i) < this.Length {
		return this.items[i]
	}

	return nil
}

func NewPyTuple(len int64) PyTuple {
	tuple := PyTuple{
		Length: len,
	}
	tuple.items = make([]PyObject, len)

	return tuple
}
