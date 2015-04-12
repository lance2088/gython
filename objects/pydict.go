package objects

import (
	"errors"
	"strings"
)

type PyDict struct {
	elements map[string]PyObject
}

func NewPyDict() PyDict {
	return PyDict{
		elements: map[string]PyObject{},
	}
}

func (this PyDict) GetType() rune { return TYPE_DICT }

func (this PyDict) String() string {
	parts := []string{}
	for key, value := range this.elements {
		parts = append(parts, "'"+key+"': "+value.String())
	}
	return "{" + strings.Join(parts, ", ") + "}"
}

func (this PyDict) Set(key string, value PyObject) {
	this.elements[key] = value
}

func (this PyDict) Exists(key string) bool {
	_, exists := this.elements[key]
	return exists
}

func (this PyDict) Get(key string) (PyObject, error) {
	value, exists := this.elements[key]
	if exists == false {
		return nil, errors.New("KeyError")
	}
	return value, nil
}
