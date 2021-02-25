package parts

type Rutherford struct {
}

func (engine *Rutherford) Mass() Kg {
	return 35
}

func (engine *Rutherford) Thrust() Newton {
	return 25000
}

func NewRutherford() Engine {
	engine := Rutherford{}
	return &engine
}
