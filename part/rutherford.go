package part

import . "github.com/ordovician/rocket/physics"

// A rocket engine with an electrical turbo pump used in the Electron rocket
// developed by Rocket Labs in New Zealand.
type Rutherford struct {
}

// The mass of the rocket engine.
func (engine Rutherford) Mass() Kg {
	return 35
}

// Thrust is the force produced when the rocket is firing.
// Think of this as similar to the horse power of a car.
func (engine Rutherford) Thrust() Newton {
	return 25000
}

// Specific impulse of the rocket engine. This is a measure of fuel efficiency.
// This of this as the milage of a car. But instead of miles rocket engine
// efficiency is thought of in terms of how much the velocity of the rocket can be altered.
func (engine Rutherford) Isp() float64 {
	return 311
}
