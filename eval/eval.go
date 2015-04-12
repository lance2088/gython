package eval

import (
	"errors"
	"fmt"

	"github.com/brettlangdon/gython/objects"
	"github.com/brettlangdon/gython/opcode"
)

func EvalFrame(frame *objects.Frame) (objects.PyObject, error) {
	var retval objects.PyObject

	idx := 0
	getNext := func() (byte, error) {
		next, err := frame.GetByteAt(idx)
		idx += 1
		return next, err
	}

	needsBreak := false
	for {
		next, err := getNext()
		if err != nil {
			break
		}

		op := int(next)
		var oparg int
		if op >= opcode.HaveArgument {
			oparg1, err := getNext()
			if err != nil {
				return nil, err
			}
			oparg2, err := getNext()
			if err != nil {
				return nil, err
			}
			oparg = (int(oparg2) << 8) + int(oparg1)
		}

		switch op {
		case opcode.Opcodes["LOAD_CONST"]:
			arg := frame.Code.Consts.GetItem(oparg)
			frame.Stack.Push(arg)
		case opcode.Opcodes["STORE_NAME"]:
			nameObj := frame.Code.Names.GetItem(oparg)

			if nameObj.GetType() == objects.TYPE_STRING {
				name := nameObj.(objects.PyString).String()
				value := frame.Stack.Pop()
				frame.Locals.Set(name, value)
			}
		case opcode.Opcodes["LOAD_FAST"]:
			value := frame.FastLocals[oparg]
			frame.Stack.Push(value)
		case opcode.Opcodes["LOAD_NAME"]:
			nameObj := frame.Code.Names.GetItem(oparg)
			if nameObj.GetType() == objects.TYPE_STRING {
				name := nameObj.(objects.PyString).String()
				value, err := frame.Locals.Get(name)
				if err != nil {
					return nil, err
				}
				frame.Stack.Push(value)
			}
		case opcode.Opcodes["PRINT_ITEM"]:
			value := frame.Stack.Pop()
			fmt.Printf("%s", value)
		case opcode.Opcodes["PRINT_NEWLINE"]:
			fmt.Printf("\r\n")
		case opcode.Opcodes["RETURN_VALUE"]:
			retval = frame.Stack.Pop()
		case opcode.Opcodes["MAKE_FUNCTION"]:
			value := frame.Stack.Pop()
			frame.Stack.Push(value)
		case opcode.Opcodes["CALL_FUNCTION"]:
			args := objects.NewPyTuple(int64(oparg))
			for i := 0; i < oparg; i++ {
				args.SetItem(i, frame.Stack.Pop())
			}
			code := frame.Stack.Pop()
			if code.GetType() != objects.TYPE_CODE {
				return nil, errors.New("Tried to execute non-code type")
			}
			newFrame := objects.NewFrame(code.(objects.PyCode), args)
			value, err := EvalFrame(newFrame)
			if err != nil {
				return nil, err
			}
			frame.Stack.Push(value)
		case opcode.Opcodes["POP_TOP"]:
			frame.Stack.Pop()
		default:
			fmt.Println("Unhandled opcode: " + opcode.Opname[op])
			needsBreak = true
		}
		if needsBreak {
			break
		}
	}

	return retval, nil
}
