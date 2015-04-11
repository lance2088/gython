package marshal

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/brettlangdon/gython/objects"
	"github.com/brettlangdon/gython/packing"
)

const SIZE32_MAX = 0x7FFFFFFF
const UNEXPECTED_TYPE = "Got unexpected type"

type Loader struct {
	buffer   *bytes.Reader
	filename string
	strings  []objects.PyString
	Code     objects.PyCode
	Magic    int32
	Modtime  int64
}

func Load(filename string) (*Loader, error) {

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	loader := &Loader{
		buffer:   bytes.NewReader(contents),
		filename: filename,
		strings:  make([]objects.PyString, 0),
	}

	err = loader.Load()
	return loader, err
}

func (this *Loader) String() string {
	output := fmt.Sprintf("Filename: %s\r\nMagic: %s\r\nModetime: %s\r\n", this.filename, this.Magic, this.Modtime)
	output += fmt.Sprintf("Strings:\r\n")
	len := len(this.strings)
	for i := 0; i < len; i++ {
		next := this.strings[i]
		output += fmt.Sprintf("\t\"%s\"\r\n", next)
	}

	output += fmt.Sprintf("Code:\r\n\t%s\r\n", this.Code)

	return output
}

func (this *Loader) Load() error {
	var err error
	this.Magic, err = this.readMagic()
	if err != nil {
		return err
	}

	this.Modtime, err = this.readModtime()
	if err != nil {
		return err
	}

	code, err := this.readObject()
	if err != nil {
		return err
	}

	if code == nil {
		return errors.New(fmt.Sprintf("Expected PyObject, got %v", code))
	}

	if code.GetType() != objects.TYPE_CODE {
		return errors.New(fmt.Sprintf("Expected TYPE_CODE, got %s", objects.TYPES[code.GetType()]))
	}
	this.Code = code.(objects.PyCode)

	return nil
}

func (this *Loader) read(length int) (buf []byte, err error) {
	buf = make([]byte, length)
	for i := 0; i < length; i++ {
		next, err := this.buffer.ReadByte()
		if err != nil {
			return buf, err
		}

		buf[i] = next
	}

	return buf, nil
}

func (this *Loader) readMagic() (int32, error) {
	next, err := this.read(4)
	if err != nil {
		return 0, nil
	}
	results, err := packing.Unpack("Hcc", next)
	if err != nil {
		return 0, nil
	}

	return results[0].(int32), nil
}

func (this *Loader) readModtime() (int64, error) {
	next, err := this.read(4)

	if err != nil {
		return 0, err
	}

	results, err := packing.Unpack("L", next)
	if err != nil {
		return 0, err
	}

	return results[0].(int64), nil
}

func (this *Loader) readObject() (obj objects.PyObject, err error) {
	chars, err := this.read(1)
	if err != nil {
		return nil, err
	}
	char := chars[0]

	switch char {
	case objects.TYPE_CODE:
		code := objects.PyCode{}
		code.ArgCount, err = this.readLong()
		if err != nil {
			return nil, err
		}

		code.NumLocals, err = this.readLong()
		if err != nil {
			return nil, err
		}

		code.StackSize, err = this.readLong()
		if err != nil {
			return nil, err
		}

		code.Flags, err = this.readLong()
		if err != nil {
			return nil, err
		}

		codeObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if codeObj.GetType() != objects.TYPE_STRING {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Code = codeObj.(objects.PyString)

		constsObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if constsObj.GetType() != objects.TYPE_TUPLE {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Consts = constsObj.(objects.PyTuple)

		namesObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if namesObj.GetType() != objects.TYPE_TUPLE {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Names = namesObj.(objects.PyTuple)

		varnamesObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if varnamesObj.GetType() != objects.TYPE_TUPLE {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Varnames = varnamesObj.(objects.PyTuple)

		freevarsObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if freevarsObj.GetType() != objects.TYPE_TUPLE {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Freevars = freevarsObj.(objects.PyTuple)

		cellvarsObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if cellvarsObj.GetType() != objects.TYPE_TUPLE {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Cellvars = cellvarsObj.(objects.PyTuple)

		filenameObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if filenameObj.GetType() != objects.TYPE_STRING {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Filename = filenameObj.(objects.PyString)

		nameObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if nameObj.GetType() != objects.TYPE_STRING {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.Name = nameObj.(objects.PyString)

		code.FirstLineNo, err = this.readLong()
		if err != nil {
			return nil, err
		}

		lnotabObj, err := this.readObject()
		if err != nil {
			return nil, err
		}
		if lnotabObj.GetType() != objects.TYPE_STRING {
			return nil, errors.New(UNEXPECTED_TYPE)
		}
		code.LNoTab = lnotabObj.(objects.PyString)

		obj = code
	case objects.TYPE_STRING, objects.TYPE_INTERNED:
		len, err := this.readLong()
		if err != nil {
			return nil, err
		}
		if len < 0 || len > SIZE32_MAX {
			return nil, errors.New("Unicode size out of range")
		}

		chars, err := this.read(int(len))
		if err != nil {
			return nil, err
		}

		str := objects.PyString{
			Length: len,
			Chars:  chars,
		}

		if char == objects.TYPE_INTERNED {
			this.strings = append(this.strings, str)
		}
		obj = str
	case objects.TYPE_TUPLE:
		len, err := this.readLong()
		if err != nil {
			return nil, err
		}
		if len < 0 || len > SIZE32_MAX {
			return nil, errors.New("Tuple size out of range")
		}

		tuple := objects.NewPyTuple(len)

		var next objects.PyObject
		for i := 0; i < int(len); i++ {
			next, err = this.readObject()
			if err != nil {
				return nil, err
			}
			tuple.SetItem(i, next)
		}
		obj = tuple
	case objects.TYPE_INT:
		num, err := this.readLong()
		if err != nil {
			return nil, err
		}
		obj = objects.PyInt{
			Number: int32(num),
		}
	case objects.TYPE_NONE:
		obj = objects.PyNone{}
	case objects.TYPE_STRINGREF:
		i, err := this.readLong()
		if err != nil {
			return nil, err
		}
		if i < 0 || int(i) >= len(this.strings) {
			return nil, errors.New("String ref out of range")
		}

		obj = this.strings[i]

	default:
		return nil, errors.New(fmt.Sprintf("Unexpected type identifier '%v'", string(char)))
	}

	return obj, nil
}

func (this *Loader) readLong() (int64, error) {
	next, err := this.read(4)
	if err != nil {
		return 0, err
	}

	results, err := packing.Unpack("L", next)
	if err != nil {
		return 0, err
	}

	return results[0].(int64), nil
}
