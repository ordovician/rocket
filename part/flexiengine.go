package part

import . "github.com/ordovician/rocket/physics"

// A rocket engine you can configure the parameters of yourself
type FlexiEngine struct {
	mass   Kg
	thrust Newton
	isp    float64
}

// The mass of the rocket engine.
func (engine *FlexiEngine) Mass() Kg {
	return engine.mass
}

// Thrust is the force produced when the rocket is firing.
// Think of this as similar to the horse power of a car.
func (engine *FlexiEngine) Thrust() Newton {
	return engine.thrust
}

// Specific impulse of the rocket engine. This is a measure of fuel efficiency.
// This of this as the milage of a car. But instead of miles rocket engine
// efficiency is thought of in terms of how much the velocity of the rocket can be altered.
func (engine *FlexiEngine) Isp() float64 {
	return engine.isp
}

// Create a configurable rocket engine
func NewFlexiEngine(mass Kg, thrust Newton, Isp float64) *FlexiEngine {
	return &FlexiEngine{mass, thrust, Isp}
}
