package parts

import . "github.com/ordovician/rockets/physics"

type Rutherford struct {
}

func (engine *Rutherford) Mass() Kg {
	return 35
}

func (engine *Rutherford) Thrust() Newton {
	return 25000
}

func (engine *Rutherford) Isp() float64 {
	return 311
}

func NewRutherford() Engine {
	engine := Rutherford{}
	return &engine
}
