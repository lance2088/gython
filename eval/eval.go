package eval

import (
	"errors"
	"fmt"

	"github.com/brettlangdon/gython/objects"
	"github.com/brettlangdon/gython/opcode"
)

func Eval(code objects.PyCode) error {
	if code.Code.Length <= 0 {
		return errors.New("No bytecodes to execute")
	}

	stack := NewStack()
	locals := map[string]objects.PyObject{}
	var retval objects.PyObject

	idx := 0
	getNext := func() (byte, error) {
		next, err := code.Code.GetIndex(idx)
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
				return err
			}
			oparg2, err := getNext()
			if err != nil {
				return err
			}
			oparg = (int(oparg2) << 8) + int(oparg1)
		}

		switch op {
		case opcode.Opcodes["LOAD_CONST"]:
			arg := code.Consts.GetItem(oparg)
			stack.Push(arg)
		case opcode.Opcodes["STORE_NAME"]:
			nameObj := code.Names.GetItem(oparg)

			if nameObj.GetType() == objects.TYPE_STRING {
				name := nameObj.(objects.PyString).String()
				value := stack.Pop()
				locals[name] = value
			}
		case opcode.Opcodes["LOAD_NAME"]:
			nameObj := code.Names.GetItem(oparg)
			if nameObj.GetType() == objects.TYPE_STRING {
				name := nameObj.(objects.PyString).String()
				value, ok := locals[name]
				if ok == false {
					return errors.New("Unknown name")
				}
				stack.Push(value)
			}
		case opcode.Opcodes["PRINT_ITEM"]:
			value := stack.Pop()
			fmt.Printf("%s", value)
		case opcode.Opcodes["PRINT_NEWLINE"]:
			fmt.Printf("\r\n")
		case opcode.Opcodes["RETURN_VALUE"]:
			retval = stack.Pop()
		case opcode.Opcodes["MAKE_FUNCTION"]:
			value := stack.Pop()
			stack.Push(value)
		default:
			needsBreak = true
		}
		if needsBreak {
			break
		}
	}

	if retval != nil {

	}
	return nil
}
