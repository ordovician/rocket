package rocket

import (
	. "github.com/ordovician/rocket/parts"
	. "github.com/ordovician/rocket/physics"
)

// A space craft is what you would think of as a space ship.
// It is the part of a multistaged rocket which will be navigating
// around in outer space.
type SpaceCraft struct {
	Payload Part
	Propulsion
}

// Create a new space craft with given payload tank and engine
func NewSpaceCraft(payload Part, tank Tank, engine Engine) *SpaceCraft {
	return &SpaceCraft{
		payload,
		Propulsion{tank, engine},
	}
}

// The mass of the space craf including engine, tank, propellant and rocket engine
func (craft *SpaceCraft) Mass() Kg {
	mass := craft.Propulsion.Mass()
	if craft.Payload != nil {
		mass += craft.Payload.Mass()
	}
	return mass
}

// SpaceCraft has no stages and thus cannot be stage separated. Thus this
// will always return nil
func (craft *SpaceCraft) StageSeparate() Rocket {
	return &NullRocket{}
}

// Number of stages for a spacecraft. A space craft is a single stage rocket
// so this will allways return 1.
func (craft *SpaceCraft) NoStages() int {
	return 1
}
