package vehicles

import . "github.com/ordovician/rocket/parts"

// A rocket is made up of some kind payload, a propellant tank and a rocket engine.
// This is a composite structure. The payload could potentially be another rocket
// to create a staged rocket.
type Rocket interface {
	Part
	Tank
	Engine

	Update(Δt float64)
	StageSeparate() Rocket
	NoStages() int
}
