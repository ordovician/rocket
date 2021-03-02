package rocket

import (
	. "github.com/ordovician/rocket/part"
	. "github.com/ordovician/rocket/physics"
)

// MultiStaged rocket means  a rocket made up of multiple stages. Thus the payload can be another rocket.
// This rocket can do stage separation.
type MultiStaged struct {
	payload Rocket
	Propulsion
}

// NewMultiStaged creates a new multi-staged rocket made up of given payload, tank and rocket engine.
func NewMultiStaged(payload Rocket, tank Tank, engine Engine) *MultiStaged {
	return &MultiStaged{
		payload,
		Propulsion{tank, engine},
	}
}

// Mass of multi-staged rocket with its engine, tank and payload
func (stage *MultiStaged) Mass() Kg {
	return stage.Propulsion.Mass() + stage.payload.Mass()
}

// StageSeparate the rocket from its payload returning the payload.
// if there is not payload, then a NullRocket will be returned instead.
func (stage *MultiStaged) StageSeparate() Rocket {
	payload := stage.payload
	stage.payload = &NullRocket{}
	return payload
}

// NoStages in a multi-stage rocket. Each stage is one rocket. A typical rocket consists
// of two stages. Something like a spaceplane would be considered single stage. In this case
// NoStages would return 1. While a NullRocket returns 0, because it isn't anything.
func (stage *MultiStaged) NoStages() int {
	return 1 + stage.payload.NoStages()
}
