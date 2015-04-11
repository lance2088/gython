package objects

const TYPE_NULL = '0'
const TYPE_NONE = 'N'
const TYPE_FALSE = 'F'
const TYPE_TRUE = 'T'
const TYPE_STOPITER = 'S'
const TYPE_ELLIPSIS = '.'
const TYPE_INT = 'i'
const TYPE_INT64 = 'I'
const TYPE_FLOAT = 'f'
const TYPE_BINARY_FLOAT = 'g'
const TYPE_COMPLEX = 'x'
const TYPE_BINARY_COMPLEX = 'y'
const TYPE_LONG = 'l'
const TYPE_STRING = 's'
const TYPE_INTERNED = 't'
const TYPE_STRINGREF = 'R'
const TYPE_TUPLE = '('
const TYPE_LIST = '['
const TYPE_DICT = '{'
const TYPE_CODE = 'c'
const TYPE_UNICODE = 'u'
const TYPE_UNKNOWN = '?'
const TYPE_SET = '<'
const TYPE_FROZENSET = '>'

var TYPES = map[rune]string{
	TYPE_NULL:           "TYPE_NULL",
	TYPE_NONE:           "TYPE_NONE",
	TYPE_FALSE:          "TYPE_FALSE",
	TYPE_TRUE:           "TYPE_TRUE",
	TYPE_STOPITER:       "TYPE_STOPITER",
	TYPE_ELLIPSIS:       "TYPE_ELLIPSIS",
	TYPE_INT:            "TYPE_INT",
	TYPE_INT64:          "TYPE_INT64",
	TYPE_FLOAT:          "TYPE_FLOAT",
	TYPE_BINARY_FLOAT:   "TYPE_BINARY_FLOAT",
	TYPE_COMPLEX:        "TYPE_COMPLEX",
	TYPE_BINARY_COMPLEX: "TYPE_BINARY_COMPLEX",
	TYPE_LONG:           "TYPE_LONG",
	TYPE_STRING:         "TYPE_STRING",
	TYPE_INTERNED:       "TYPE_INTERNED",
	TYPE_STRINGREF:      "TYPE_STRINGREF",
	TYPE_TUPLE:          "TYPE_TUPLE",
	TYPE_LIST:           "TYPE_LIST",
	TYPE_DICT:           "TYPE_DICT",
	TYPE_CODE:           "TYPE_CODE",
	TYPE_UNICODE:        "TYPE_UNICODE",
	TYPE_UNKNOWN:        "TYPE_UNKNOWN",
	TYPE_SET:            "TYPE_SET",
	TYPE_FROZENSET:      "TYPE_FROZENSET",
}
