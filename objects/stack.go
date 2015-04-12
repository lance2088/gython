package objects

type Stack struct {
	elements []PyObject
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]PyObject, 0),
	}
}

func (this *Stack) Len() int {
	return len(this.elements)
}

func (this *Stack) Push(elm PyObject) {
	this.elements = append(this.elements, elm)
}

func (this *Stack) Pop() PyObject {
	value := this.elements[this.Len()-1]
	this.elements = this.elements[:this.Len()-1]

	return value
}
