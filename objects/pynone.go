package objects

type PyNone struct{}

func (this PyNone) GetType() rune { return TYPE_NONE }

func (this PyNone) String() string {
	return "None"
}
