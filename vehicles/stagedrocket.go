package vehicles

import (
	. "github.com/ordovician/rocket/parts"
	. "github.com/ordovician/rocket/physics"
)

// A multi-staged rocket. Meaning a rocket where the payload is another rocket.
// This rocket can do stage separation.
type StagedRocket struct {
	payload Rocket
	Propulsion
}

func NewStagedRocket(payload Rocket, tank Tank, engine Engine) *StagedRocket {
	return &StagedRocket{
		payload,
		Propulsion{tank, engine},
	}
}

// Mass of multi-staged rocket with its engine, tank and payload
func (stage *StagedRocket) Mass() Kg {
	return stage.Propulsion.Mass() + stage.payload.Mass()
}

// Detach payload from rocket. Payload is returned
// The payload of the staged rocket is set to a NullRocket to indicate
// there is no payload anymore
func (stage *StagedRocket) StageSeparate() Rocket {
	payload := stage.payload
	stage.payload = &NullRocket{}
	return payload
}

// Number of stages in a multistage rocket.
// Will at minumum be 1. If there is no payload (NullRocket) then thee
// number of stages will be 1
func (stage *StagedRocket) NoStages() int {
	return 1 + stage.payload.NoStages()
}
