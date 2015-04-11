package dis

import (
	"fmt"
	"strconv"

	"github.com/brettlangdon/gython/objects"
	"github.com/brettlangdon/gython/opcode"
)

func Disassemble(code objects.PyCode) {
	lines := code.Linestarts()
	codeStr := code.Code.String()
	n := len(codeStr)
	extendedArg := 0
	for i := 0; i < n; {
		c := codeStr[i]
		op := int(c)

		if lineNo, ok := lines[i]; ok {
			if i > 0 {
				fmt.Println("")
			}
			fmt.Printf("%3s", strconv.FormatInt(int64(lineNo), 10))
		} else {
			fmt.Printf("   ")
		}

		fmt.Printf("  ")
		fmt.Printf("%-4s", strconv.FormatInt(int64(i), 10))
		fmt.Printf("%-20s", opcode.Opname[op])

		i += 1
		if op >= opcode.HaveArgument {
			oparg := int(codeStr[i]) + int(codeStr[i+1])*256 + extendedArg
			extendedArg = 0
			i += 2

			if op == opcode.ExtendedArg {
				extendedArg = oparg * 65536
			}

			fmt.Printf("%-5s", strconv.FormatInt(int64(oparg), 10))

			if _, ok := opcode.HasConst[op]; ok {
				fmt.Printf("(" + code.Consts.GetItem(oparg).String() + ")")
			} else if _, ok := opcode.HasName[op]; ok {
				fmt.Printf("(" + code.Names.GetItem(oparg).String() + ")")
			} else if _, ok := opcode.HasLocal[op]; ok {
				fmt.Printf("(" + code.Varnames.GetItem(oparg).String() + ")")
			}
		}

		fmt.Printf("\r\n")
	}
}
