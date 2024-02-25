package pybool

type PyBool bool

func New(b bool) PyBool {
	return PyBool(b)
}

func (pb PyBool) Bool() bool {
	return bool(pb)
}

func (pb PyBool) String() string {
	if bool(pb) {
		return "True"
	}
	return "False"
}
