package eval

import "github.com/brettlangdon/gython/objects"

type Stack struct {
	elements []objects.PyObject
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]objects.PyObject, 0),
	}
}

func (this *Stack) Len() int {
	return len(this.elements)
}

func (this *Stack) Push(elm objects.PyObject) {
	this.elements = append(this.elements, elm)
}

func (this *Stack) Pop() objects.PyObject {
	value := this.elements[this.Len()-1]
	this.elements = this.elements[:this.Len()-1]

	return value
}
