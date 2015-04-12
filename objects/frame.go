package objects

type Frame struct {
	Args          PyTuple
	Builtins      PyDict
	Code          PyCode
	FastLocals    []PyObject
	Globals       PyDict
	Locals        PyDict
	PreviousFrame *Frame
	Stack         *Stack
}

func NewFrame(code PyCode, args PyTuple) *Frame {
	frame := &Frame{
		Args:          args,
		Builtins:      NewPyDict(),
		Code:          code,
		FastLocals:    []PyObject{},
		Globals:       NewPyDict(),
		Locals:        NewPyDict(),
		PreviousFrame: nil,
		Stack:         NewStack(),
	}

	// TODO: finish implementing function arguments
	//   https://github.com/python/cpython/blob/2.7/Python/ceval.c#L3055-L3265
	if frame.Code.ArgCount > 0 || frame.Code.Flags&(CO_VARARGS|CO_VARKEYWORDS) > 0 {
		for i := 0; i < int(frame.Args.Length); i++ {
			frame.FastLocals = append(frame.FastLocals, frame.Args.GetItem(i))
		}
	}

	return frame
}

func (this *Frame) GetByteAt(idx int) (byte, error) {
	return this.Code.Code.GetIndex(idx)
}
