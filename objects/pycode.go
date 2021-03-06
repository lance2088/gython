package objects

import (
	"fmt"
	"strconv"
)

const CO_OPTIMIZED = 0x0001
const CO_NEWLOCALS = 0x0002
const CO_VARARGS = 0x0004
const CO_VARKEYWORDS = 0x0008
const CO_NESTED = 0x0010
const CO_GENERATOR = 0x0020
const CO_NOFREE = 0x0040

type PyCode struct {
	ArgCount    int64
	Cellvars    PyTuple
	Code        PyString
	Consts      PyTuple
	Filename    PyString
	FirstLineNo int64
	Flags       int64
	Freevars    PyTuple
	LNoTab      PyString
	Name        PyString
	Names       PyTuple
	NumLocals   int64
	StackSize   int64
	Varnames    PyTuple
}

func (this PyCode) GetType() rune { return TYPE_CODE }

func (this PyCode) String() string {
	return fmt.Sprintf("<code object %s, file %s, line %s>", this.Name, this.Filename, strconv.FormatInt(this.FirstLineNo, 10))
}

func (this PyCode) Linestarts() map[int]int {
	lines := make(map[int]int)
	byteIncrements := []int{}
	lineIncrements := []int{}

	lnotab := this.LNoTab.String()
	for i := 0; i < len(lnotab); i += 2 {
		byteIncrements = append(byteIncrements, int(lnotab[i]))
	}

	for i := 1; i < len(lnotab); i += 2 {
		lineIncrements = append(lineIncrements, int(lnotab[i]))
	}

	lastLineNo := -1
	lineNo := int(this.FirstLineNo)
	addr := 0

	for i := 0; i < len(byteIncrements); i++ {
		if byteIncrements[i] > 0 {
			if lineNo != lastLineNo {
				lines[addr] = lineNo
			}
			addr += byteIncrements[i]
		}
		lineNo += lineIncrements[i]
	}

	if lineNo != lastLineNo {
		lines[addr] = lineNo
	}

	return lines
}
