package engine

import (
	"encoding/json"

	. "github.com/ordovician/rocket/physics"
)

// A rocket engine you can configure the parameters of yourself
type CustomEngine struct {
	mass   Kg
	thrust Newton
	isp    float64
}

// MarshalJSON implements the Marshaler interface to allow custom serialization
// of a data structure. In this case CustomEngine has private fields and
// thus cannot be automatically serialized and deserialized
func (engine *CustomEngine) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(struct {
		Mass   Kg
		Thrust Newton
		Isp    float64
	}{
		Mass:   engine.mass,
		Thrust: engine.thrust,
		Isp:    engine.isp,
	})

	if err != nil {
		return nil, err
	}
	return data, nil
}

// The mass of the rocket engine.
func (engine *CustomEngine) Mass() Kg {
	return engine.mass
}

// Thrust is the force produced when the rocket is firing.
// Think of this as similar to the horse power of a car.
func (engine *CustomEngine) Thrust() Newton {
	return engine.thrust
}

// Specific impulse of the rocket engine. This is a measure of fuel efficiency.
// This of this as the milage of a car. But instead of miles rocket engine
// efficiency is thought of in terms of how much the velocity of the rocket can be altered.
func (engine *CustomEngine) Isp() float64 {
	return engine.isp
}

// Create a configurable rocket engine
func NewCustomEngine(mass Kg, thrust Newton, Isp float64) *CustomEngine {
	return &CustomEngine{mass, thrust, Isp}
}
