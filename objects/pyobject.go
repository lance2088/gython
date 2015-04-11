package objects

type PyObject interface {
	GetType() rune
	String() string
}
