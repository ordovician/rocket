package part

import . "github.com/ordovician/rocket/physics"

// Something that has mass
type Part interface {
	Mass() Kg
}

// A tank to hold rocket propellant
type Tank interface {
	Part
	Propellant() Kg
	IsEmpty() bool
	Consume(amount Kg) Kg
	Refill()
}

// A rocket engine
type Engine interface {
	Part
	Thrust() Newton // Power of engine. How much it pushes
	Isp() float64   // Specific impulse. Fuel efficiency of engine
}