package tank

import (
	. "github.com/ordovician/rocket/physics"
)

// A tank which can contain nothing
type NullTank struct {
}

func (tank *NullTank) Propellant() Kg {
	return 0
}

// Is the tank empty?
func (tank *NullTank) IsEmpty() bool {
	return true
}

// Consume given amount of propellant. Returns how much was consumed
func (tank *NullTank) Consume(amount Kg) Kg {
	return 0
}

// Fill tank up to the max
func (tank *NullTank) Refill() {
}

// A null tank has no mass because it isn't a tank. It is a nothing tank
func (tank *NullTank) Mass() Kg {
	return 0
}
