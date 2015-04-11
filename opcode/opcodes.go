package opcode

import "fmt"

var ComparisonOperators = []string{
	"<", "<=", "==", "!=", ">", ">=", "in", "not in", "is",
	"is not", "exception match", "BAD",
}

var HasConst = make(map[int]bool)
var HasName = make(map[int]bool)
var HasJrel = make(map[int]bool)
var HasJabs = make(map[int]bool)
var HasLocal = make(map[int]bool)
var HasCompare = make(map[int]bool)
var HasFree = make(map[int]bool)
var HaveArgument = 90
var ExtendedArg = 145

var Opcodes = make(map[string]int)
var Opname = []string{}

func opCode(name string, code int) {
	Opcodes[name] = code
	Opname[code] = name
}

func nameCode(name string, code int) {
	opCode(name, code)
	HasName[code] = true
}

func jrelCode(name string, code int) {
	opCode(name, code)
	HasJrel[code] = true
}

func jabsCode(name string, code int) {
	opCode(name, code)
	HasJabs[code] = true
}

func localCode(name string, code int) {
	opCode(name, code)
	HasLocal[code] = true
}

func compareCode(name string, code int) {
	opCode(name, code)
	HasCompare[code] = true
}

func freeCode(name string, code int) {
	opCode(name, code)
	HasFree[code] = true
}

func constCode(name string, code int) {
	opCode(name, code)
	HasConst[code] = true
}

func defineOpcodes() {
	opCode("STOP_CODE", 0)
	opCode("POP_TOP", 1)
	opCode("ROT_TWO", 2)
	opCode("ROT_THREE", 3)
	opCode("DUP_TOP", 4)
	opCode("ROT_FOUR", 5)

	opCode("NOP", 9)
	opCode("UNARY_POSITIVE", 10)
	opCode("UNARY_NEGATIVE", 11)
	opCode("UNARY_NOT", 12)
	opCode("UNARY_CONVERT", 13)

	opCode("UNARY_INVERT", 15)

	opCode("BINARY_POWER", 19)
	opCode("BINARY_MULTIPLY", 20)
	opCode("BINARY_DIVIDE", 21)
	opCode("BINARY_MODULO", 22)
	opCode("BINARY_ADD", 23)
	opCode("BINARY_SUBTRACT", 24)
	opCode("BINARY_SUBSCR", 25)
	opCode("BINARY_FLOOR_DIVIDE", 26)
	opCode("BINARY_TRUE_DIVIDE", 27)
	opCode("INPLACE_FLOOR_DIVIDE", 28)
	opCode("INPLACE_TRUE_DIVIDE", 29)
	opCode("SLICE+0", 30)
	opCode("SLICE+1", 31)
	opCode("SLICE+2", 32)
	opCode("SLICE+3", 33)

	opCode("STORE_SLICE+0", 40)
	opCode("STORE_SLICE+1", 41)
	opCode("STORE_SLICE+2", 42)
	opCode("STORE_SLICE+3", 43)

	opCode("DELETE_SLICE+0", 50)
	opCode("DELETE_SLICE+1", 51)
	opCode("DELETE_SLICE+2", 52)
	opCode("DELETE_SLICE+3", 53)

	opCode("STORE_MAP", 54)
	opCode("INPLACE_ADD", 55)
	opCode("INPLACE_SUBTRACT", 56)
	opCode("INPLACE_MULTIPLY", 57)
	opCode("INPLACE_DIVIDE", 58)
	opCode("INPLACE_MODULO", 59)
	opCode("STORE_SUBSCR", 60)
	opCode("DELETE_SUBSCR", 61)
	opCode("BINARY_LSHIFT", 62)
	opCode("BINARY_RSHIFT", 63)
	opCode("BINARY_AND", 64)
	opCode("BINARY_XOR", 65)
	opCode("BINARY_OR", 66)
	opCode("INPLACE_POWER", 67)
	opCode("GET_ITER", 68)

	opCode("PRINT_EXPR", 70)
	opCode("PRINT_ITEM", 71)
	opCode("PRINT_NEWLINE", 72)
	opCode("PRINT_ITEM_TO", 73)
	opCode("PRINT_NEWLINE_TO", 74)
	opCode("INPLACE_LSHIFT", 75)
	opCode("INPLACE_RSHIFT", 76)
	opCode("INPLACE_AND", 77)
	opCode("INPLACE_XOR", 78)
	opCode("INPLACE_OR", 79)
	opCode("BREAK_LOOP", 80)
	opCode("WITH_CLEANUP", 81)
	opCode("LOAD_LOCALS", 82)
	opCode("RETURN_VALUE", 83)
	opCode("IMPORT_STAR", 84)
	opCode("EXEC_STMT", 85)
	opCode("YIELD_VALUE", 86)
	opCode("POP_BLOCK", 87)
	opCode("END_FINALLY", 88)
	opCode("BUILD_CLASS", 89)

	nameCode("STORE_NAME", 90)
	nameCode("DELETE_NAME", 91)
	opCode("UNPACK_SEQUENCE", 92)
	jrelCode("FOR_ITER", 93)
	opCode("LIST_APPEND", 94)
	nameCode("STORE_ATTR", 95)
	nameCode("DELETE_ATTR", 96)
	nameCode("STORE_GLOBAL", 97)
	nameCode("DELETE_GLOBAL", 98)
	opCode("DUP_TOPX", 99)
	constCode("LOAD_CONST", 100)
	nameCode("LOAD_NAME", 101)
	opCode("BUILD_TUPLE", 102)
	opCode("BUILD_LIST", 103)
	opCode("BUILD_SET", 104)
	opCode("BUILD_MAP", 105)
	nameCode("LOAD_ATTR", 106)
	compareCode("COMPARE_OP", 107)
	nameCode("IMPORT_NAME", 108)
	nameCode("IMPORT_FROM", 109)
	jrelCode("JUMP_FORWARD", 110)
	jabsCode("JUMP_IF_FALSE_OR_POP", 111)
	jabsCode("JUMP_IF_TRUE_OR_POP", 112)
	jabsCode("JUMP_ABSOLUTE", 113)
	jabsCode("POP_JUMP_IF_FALSE", 114)
	jabsCode("POP_JUMP_IF_TRUE", 115)

	nameCode("LOAD_GLOBAL", 116)

	jabsCode("CONTINUE_LOOP", 119)
	jrelCode("SETUP_LOOP", 120)
	jrelCode("SETUP_EXCEPT", 121)
	jrelCode("SETUP_FINALLY", 122)

	localCode("LOAD_FAST", 124)
	localCode("STORE_FAST", 125)
	localCode("DELETE_FAST", 126)

	opCode("RAISE_VARARGS", 130)
	opCode("CALL_FUNCTION", 131)
	opCode("MAKE_FUNCTION", 132)
	opCode("BUILD_SLICE", 133)
	opCode("MAKE_CLOSURE", 134)
	freeCode("LOAD_CLOSURE", 135)
	freeCode("LOAD_DEREF", 136)
	freeCode("STORE_DEREF", 137)

	opCode("CALL_FUNCTION_VAR", 140)
	opCode("CALL_FUNCTION_KW", 141)
	opCode("CALL_FUNCTION_VAR_KW", 142)

	jrelCode("SETUP_WITH", 143)

	opCode("EXTENDED_ARG", 145)
	opCode("SET_ADD", 146)
	opCode("MAP_ADD", 147)

}

func init() {
	for i := 0; i < 256; i++ {
		Opname = append(Opname, fmt.Sprintf("<%d>", i))
	}
	defineOpcodes()
}
