package engine

import . "github.com/ordovician/rocket/physics"

// A rocket engine
type Engine interface {
	Mass() Kg       // Mass of whole engine
	Thrust() Newton // Power of engine. How much it pushes
	Isp() float64   // Specific impulse. Fuel efficiency of engine
}
