package engine

import . "github.com/ordovician/rocket/physics"

// A rocket engine with an electrical turbo pump used in
// the Falcon 9 rocket
// developed by SpaceX.
type Merlin struct {
}

// The mass of the rocket engine.
func (engine Merlin) Mass() Kg {
	return 470
}

// Thrust is the force produced when the rocket is firing.
// Think of this as similar to the horse power of a car.
func (engine Merlin) Thrust() Newton {
	return 845e3
}

// Specific impulse of the rocket engine. This is a measure of fuel efficiency.
// This of this as the milage of a car. But instead of miles rocket engine
// efficiency is thought of in terms of how much the velocity of the rocket can be altered.
func (engine Merlin) Isp() float64 {
	return 282
}
