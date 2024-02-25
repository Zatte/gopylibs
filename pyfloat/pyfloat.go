package pybool

import (
	"fmt"
	"math/big"
)

type PyFloat float64

func New(f float64) PyFloat {
	return PyFloat(f)
}

func (pf PyFloat) Float() float64 {
	return float64(pf)
}

func (pf PyFloat) String() string {
	return fmt.Sprintf("%f", float64(pf))
}

func (pf PyFloat) AsIntegerRatio() (int, int) {
	r, _ := big.NewFloat(float64(pf)).Rat(nil)
	return int(r.Num().Int64()), int(r.Denom().Int64())
}

func (pf PyFloat) IsInteger() bool {
	return pf == PyFloat(int(pf))
}

func (pf PyFloat) Hex() string {
	return fmt.Sprintf("%#x", float64(pf))
}
